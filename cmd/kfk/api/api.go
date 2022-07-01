package main

import (
	"encoding/json"
	oidc "github.com/coreos/go-oidc"
	"github.com/gin-gonic/gin"
	"golang.org/x/net/context"
	"golang.org/x/oauth2"
	"log"
	"net/http"
)

type OidcConfig struct {
	ClientID        string
	ClientSecret    string
	AppNonce        string
	Provider        string
	IDTokenVerifier *oidc.IDTokenVerifier
	Oauth2Config    oauth2.Config
	State           string
}

func initOidc() *OidcConfig {

	o := &OidcConfig{
		ClientID:     "test_1",
		ClientSecret: "secret",
		AppNonce:     "a super secret nonce",
		Provider:     "http://172.19.0.128:4444/",
	}
	ctx := context.Background()

	provider, err := oidc.NewProvider(ctx, o.Provider)
	if err != nil {
		log.Fatal(err)
	}

	oidcConfig := &oidc.Config{
		ClientID: clientID,
	}
	// Use the nonce source to create a custom ID Token verifier.
	nonceEnabledVerifier := provider.Verifier(oidcConfig)

	o.IDTokenVerifier = nonceEnabledVerifier

	config := oauth2.Config{
		ClientID:     clientID,
		ClientSecret: clientSecret,
		Endpoint:     provider.Endpoint(),
		RedirectURL:  "http://localhost:8080/api/login/callback",
		Scopes:       []string{oidc.ScopeOpenID, "profile", "email"},
	}

	o.Oauth2Config = config
	state, _ := uuid.GenerateUUID() // Don't do this in production.
	o.State = state
	return o
}

func AuthLogin(c *gin.Context) {
	o := initOidc()
	authUrl := o.Oauth2Config.AuthCodeURL(o.State, oidc.Nonce(o.AppNonce))
	c.JSON(http.StatusOK, authUrl)
}

func CallBack(c *gin.Context) {
	o := initOidc()
	state, _ := c.GetQuery("state")
	ctx := context.Background()
	code, _ := c.GetQuery("code")
	log.Println(code)
	oauth2Token, err := o.Oauth2Config.Exchange(ctx, code)
	if err != nil {
		c.String(http.StatusInternalServerError, "Failed to exchange token: "+err.Error())
		return
	}
	rawIDToken, ok := oauth2Token.Extra("id_token").(string)
	if !ok {
		c.String(http.StatusInternalServerError, "No id_token field in oauth2 token.")
		return
	}
	idToken, err := o.IDTokenVerifier.Verify(ctx, rawIDToken)
	if err != nil {
		c.String(http.StatusInternalServerError, "Failed to verify ID Token: "+err.Error())
		return
	}
	if idToken.Nonce != o.AppNonce {
		c.String(http.StatusInternalServerError, "Invalid ID Token nonce")
		return
	}
	resp := struct {
		OAuth2Token   *oauth2.Token
		IDTokenClaims *json.RawMessage // ID Token payload is just JSON.
	}{oauth2Token, new(json.RawMessage)}
	if err := idToken.Claims(&resp.IDTokenClaims); err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		return
	}
	data, err := json.MarshalIndent(resp, "", "    ")
	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		return
	}
	c.SetCookie("token", state, 30, "/", "", false, true)

	log.Println(string(data))
	//c.String(http.StatusOK,string(data))
	//c.String(http.StatusOK,string(data))
	c.Redirect(http.StatusFound, "/")
}
