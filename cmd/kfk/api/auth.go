/*
This is an example application to demonstrate verifying an ID Token with a nonce.
*/
package main

import (
	"encoding/json"
	oidc "github.com/coreos/go-oidc"
	"golang.org/x/net/context"
	"golang.org/x/oauth2"
	"log"
	"net/http"
)

var (
	clientID     = "test_1"
	clientSecret = "secret"
)

const appNonce = "a super secret nonce"

func Auth() {
	ctx := context.Background()

	provider, err := oidc.NewProvider(ctx, "http://172.19.0.128:4444/")
	if err != nil {
		log.Fatal(err)
	}

	oidcConfig := &oidc.Config{
		ClientID: clientID,
	}
	// Use the nonce source to create a custom ID Token verifier.
	nonceEnabledVerifier := provider.Verifier(oidcConfig)

	config := oauth2.Config{
		ClientID:     clientID,
		ClientSecret: clientSecret,
		Endpoint:     provider.Endpoint(),
		RedirectURL:  "http://localhost:8080/api/login/callback",
		Scopes:       []string{oidc.ScopeOpenID, "profile", "email"},
	}

	state, _ := uuid.GenerateUUID() // Don't do this in production.

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		http.Redirect(w, r, config.AuthCodeURL(state, oidc.Nonce(appNonce)), http.StatusFound)
	})

	http.HandleFunc("/api/login/callback", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")

		if r.URL.Query().Get("state") != state {
			http.Error(w, "state did not match", http.StatusBadRequest)
			return
		}

		oauth2Token, err := config.Exchange(ctx, r.URL.Query().Get("code"))
		if err != nil {
			http.Error(w, "Failed to exchange token: "+err.Error(), http.StatusInternalServerError)
			return
		}
		//log.Printf("oauth2Token: %s\n", oauth2Token)
		rawIDToken, ok := oauth2Token.Extra("id_token").(string)
		if !ok {
			http.Error(w, "No id_token field in oauth2 token.", http.StatusInternalServerError)
			return
		}
		//log.Printf("rawIDToken: %s\n", rawIDToken)

		// Verify the ID Token signature and nonce.
		idToken, err := nonceEnabledVerifier.Verify(ctx, rawIDToken)
		if err != nil {
			http.Error(w, "Failed to verify ID Token: "+err.Error(), http.StatusInternalServerError)
			return
		}
		//log.Printf("idToken: %s\n", idToken)

		if idToken.Nonce != appNonce {
			http.Error(w, "Invalid ID Token nonce", http.StatusInternalServerError)
			return
		}

		resp := struct {
			OAuth2Token   *oauth2.Token
			IDTokenClaims *json.RawMessage // ID Token payload is just JSON.
		}{oauth2Token, new(json.RawMessage)}

		if err := idToken.Claims(&resp.IDTokenClaims); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		data, err := json.MarshalIndent(resp, "", "    ")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		//r.AddCookie(&http.Cookie{
		//	Name:       "token",
		//	Value:      resp.OAuth2Token.AccessToken,
		//	Path:       "/",
		//	Domain:     "localhost",
		//	Expires:     resp.OAuth2Token.Expiry,
		//	Secure:     false,
		//	HttpOnly:   true,
		//})

		//r.Header.Add("token",resp.OAuth2Token.AccessToken)
		//r.Header.Add("expires", resp.OAuth2Token.Expiry.String())
		//w.Header().Set("token", resp.OAuth2Token.AccessToken)
		//log.Println(data)
		http.Redirect(w, r, "http://127.0.0.1:8090", http.StatusFound)
		w.Write(data)

		return
	})

	log.Printf("listening on http://%s/", "127.0.0.1:8080")
	log.Fatal(http.ListenAndServe("127.0.0.1:8080", nil))
}
