package main

import (
	"crypto"
	"crypto/ecdsa"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"errors"
	"fmt"
	certutil "k8s.io/client-go/util/cert"
	"k8s.io/client-go/util/keyutil"
	"k8s.io/kube-openapi/pkg/util/sets"
	"net"
	"path/filepath"
)

const caCertData = `-----BEGIN CERTIFICATE-----
MIIE1TCCA72gAwIBAgIJAKgBTymfuxijMA0GCSqGSIb3DQEBCwUAMGIxCzAJBgNV
BAYTAkNOMRAwDgYDVQQIDAdCZWlqaW5nMSEwHwYDVQQKDBjDpMK4wpzDpsKWwrnD
pcKbwr3DpMK/wqExETAPBgNVBAsMCGJvbmNsb3VkMQswCQYDVQQDDAJDQTAeFw0y
MjEwMTkwOTEzMTZaFw0zMjA3MTgwOTEzMTZaMBMxETAPBgNVBAoMCEJvbmNsb3Vk
MIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEAtWqKzgXD6rBP9gwwco62
VtnI0znhj/U5SeNQY1cLi/UHNhFS3vOOsZg009A+xCKY94deRwCVCg4+JcP7knox
U89KOWqCov9aW8T0FdY1v1iczEwecnZzkviaiYubvboMHdIZwTdcnJtqB+uImDKx
nTQcZL4xEui9pz5FU9+WSMrtGDy2LhNP7AXVc8fLYpN+N25oU3s4q7gNzu5Y0RmV
X4/NRl3GdX6X12PEudb/CtlODJ7tiRga9Rq/6JNzrea/3ZRNvewQeU5s9x97xEeU
x/aACmnaSkC5wMFT0+/0mesCFjqDD4Y61qijhudZbQHDsVl0ymMuja9FqFbET49N
cQIDAQABo4IB2zCCAdcwDgYDVR0PAQH/BAQDAgWgMBMGA1UdJQQMMAoGCCsGAQUF
BwMBMAwGA1UdEwEB/wQCMAAwHwYDVR0jBBgwFoAUiQSn6K/WqpFw0uqQRVzGXyO5
IyUwggF/BgNVHREEggF2MIIBcoIGZ2NyLmlvggdnaGNyLmlvggdxdWF5Lmlvggpr
OHMuZ2NyLmlvgg9yZWdpc3RyeS5rOHMuaW+CEWRvY2tlci5lbGFzdGljLmNvghNy
ZWdpc3RyeS5naXRsYWIuY29tghVyZWdpc3RyeS5ib25jbG91ZC5jb22CE2hhcmJv
ci5ib25jbG91ZC5jb22CFGh1Yi1taXJyb3IuYy4xNjMuY29tgiByZWdpc3RyeS5j
bi1iZWlqaW5nLmFsaXl1bmNzLmNvbYISbWlycm9ycy5hbGl5dW4uY29tgg1pc28u
eXVuaW9uLmNugiZjb250YWluZXJuZXR3b3JraW5nLnBlazNiLnFpbmdzdG9yLmNv
bYIla3ViZXJuZXRlcy1yZWxlYXNlLnBlazNiLnFpbmdzdG9yLmNvbYIVcmVnaXN0
cnkuYWxpeXVuY3MuY29tghRyZWdpc3RyeS0xLmRvY2tlci5pb4IOeXVtLm9yYWNs
ZS5jb20wDQYJKoZIhvcNAQELBQADggEBAD3z/a1qLFChNQN9ZaFyF9MwlOYSchPD
WxfSVUMlVFe2BEUO5wMMQrI46jyS1PP4dmJ9eZl3FdfXq9SFS/DMnWWlmVs9xRAR
7KcckzvROUhqR9p5bDTjfk+1m8Da+XWs9bAfMFUjQv+5BwTtLLDkHC+HP5opctuR
B9w6CbrHPc6ISh/VmRvahrBm6s22crofnjHt/it8i/9va2avoBbYw45+OmEqDwmv
bAvR3V9rrFkKDXA9Uf7tboPQMG17Ytv64+CWDR7FGbxmIEXL1r3yyYkKP59rVx+A
BIkt06L6T8FTDrNg8sODn532d7fope3V38Nez79Z6HDqAT5w94Jpc5w=
-----END CERTIFICATE-----`
const caKeyData = `-----BEGIN RSA PRIVATE KEY-----
MIIEowIBAAKCAQEAyY0k8fKTYDGPf88V4UIy7AhFnYJrH/aFe9ffBE0Cr8AukabD
bz+Dbb46jkgdu3Cjsr/DdyjJ/RZnfvT7DKDWZW9zpYxvU8h9MkwyPHuOs2Ki+ZQp
UEysfPRAmuJp0JS0VQcvNSh6uufIRId/LPv5XJTVT1C9TjAnal6891+mSp8MQL3o
dah9p7IB/VIp5765HBQaquFC++CcAbY9IJlWQni8oRU6qK1X1STW710q5brt7Fij
CrIiFGMw1EgBja3TO7Qb0yaTu7ZhZ7vMCBwymJl56h3bg8sJlm3k1ngGaR/YXcOA
zvQlp77msDl1PsOesNiXXGXTAtkAL4hBFAOpZwIDAQABAoIBAFORPyFGR87ZnbsL
fPHyBbUR1PNy0MHm7/+iSOi6mYOPdn+TmnK07eSBvDukMRe3o1gI5J2yftV+BZmB
L9pTkbFrHZMlgD9H4IkLSMUFIezE1/gNc3fE4rvIdkUB3YfLMF4U5YDv8LswQTwQ
xi07qG+3kh4ZxvP1SVJ0R3t14DgNr5nP8kaDOeitMySutQugWJplbRd+1X4+xFIM
WP84qIt31oRwIJejWj/Z/glrszPqFuBp11mjgU4uYd4rqNP/ZpWXfVeB5Dg0dTvh
gBv0gCzUQN8tbYnucherLdO8IDFpyyDd+mN9ENPGgjwCd6am5kcTHGZXx9cI5hnC
+fw7fekCgYEA5/cRo6vmmmRmTSCjJgtRfph0K+xnfemsEJXYQWeEszOD7/XsR7H3
l5AlKzu6PzfdXlSyd2PCrAUe3b4xIeLzstBzN92fDHn4e7CiE0UsVO32jZP8G/TE
cRcu2BehnXUASsrmtcxCiNMDkXZjgPR/lO5IDy4FjtPjtc9y/s6U4QsCgYEA3m9X
wq2cnSR1mpjaX6XYsW4PPQ2YS1l7o/SnyugkcDiL/8aKg41CipE2mMyOyCV2uzxV
N5MfyFi8VvZrhb7eA/7aeWDG9pqkeZrPtaOH2JqzhqjiU/dWJbpaEJviuVuWCYyi
VSN3MHYcwa+0xrp8JDvNH8WmMpO9Aeu0CQueypUCgYBZ9ldP+9Y2oKOQXA0KLy+P
An9jnY7RpXOHByZUz1oGyf7sbJsTfzEABfZ0Wvizle3zrLN+XCFe56l95EpX0xYi
jndw/jG1/APjrLBe+t/jnFqXtAH9saMSHSScyCV01LClUSXC8hIH0Ja8roaOt8RX
NUabJIUhTIous+LscaAJ2wKBgHS2z8Hi/w6llt/r3Inbp/xR18UdYRAIgAvj0Ddz
38rSoQMw1nV4pbW3xIIgs7rpjYdpfP2QQVkK1qh63KhtImnOTCzsTvoO4sa8KMkS
abGKWzEJZNjSK23YfnHAmhLQr8WK+ZLa7SuMjkJDRAQSzhjlGBjXyQE47DAZ0Xn9
kvCRAoGBAI2uSw3LUoFvRuv2jvK4y5fUqOgMO4LKfbJvsn3kVTKcFhC16iRmWkAl
VD9vb/KZpQKP1wQWU9HLMLvvJgtCgHCnJpB5EyQqwLk++tC3/aba9rMTeXKQe/69
d+PHGKUJQnRJUi8+Ka8YaudCtO6hXZqmILM+DPezBaxh4Ccf8E6n
-----END RSA PRIVATE KEY-----`

const (
	// CertificateBlockType is a possible value for pem.Block.Type.
	CertificateBlockType = "CERTIFICATE"
	// CertificateRequestBlockType is a possible value for pem.Block.Type.
	CertificateRequestBlockType = "CERTIFICATE REQUEST"
)

const (
	// ECPrivateKeyBlockType is a possible value for pem.Block.Type.
	ECPrivateKeyBlockType = "EC PRIVATE KEY"
	// RSAPrivateKeyBlockType is a possible value for pem.Block.Type.
	RSAPrivateKeyBlockType = "RSA PRIVATE KEY"
	// PrivateKeyBlockType is a possible value for pem.Block.Type.
	PrivateKeyBlockType = "PRIVATE KEY"
	// PublicKeyBlockType is a possible value for pem.Block.Type.
	PublicKeyBlockType = "PUBLIC KEY"
)

// TryLoadCertAndKeyFromDisk tries to load a cert and a key from the disk and validates that they are valid
func TryLoadCertAndKey() (*x509.Certificate, crypto.Signer, error) {
	cert, err := TryLoadCert()
	if err != nil {
		return nil, nil, err
	}

	key, err := TryToLoadKey()
	if err != nil {
		return nil, nil, err
	}

	return cert, key, nil
}

// TryLoadCertFromDisk tries to load the cert from the disk
func TryLoadCert() (*x509.Certificate, error) {

	certs, err := ParseCertsPEM([]byte(caCertData))

	cert := certs[0]

	return cert, err
}

// TryLoadKeyFromDisk tries to load the key from the disk and validates that it is valid
func TryToLoadKey() (crypto.Signer, error) {

	// Parse the private key from a file
	privKey, err := ParsePrivateKeyPEM([]byte(caKeyData))
	if err != nil {
		return nil, err
	}

	// Allow RSA and ECDSA formats only
	var key crypto.Signer
	switch k := privKey.(type) {
	case *rsa.PrivateKey:
		key = k
	case *ecdsa.PrivateKey:
		key = k
	default:
		return nil, err
	}

	return key, nil
}

// ParseCertsPEM returns the x509.Certificates contained in the given PEM-encoded byte array
// Returns an error if a certificate could not be parsed, or if the data does not contain any certificates
func ParseCertsPEM(pemCerts []byte) ([]*x509.Certificate, error) {
	ok := false
	certs := []*x509.Certificate{}
	for len(pemCerts) > 0 {
		var block *pem.Block
		block, pemCerts = pem.Decode(pemCerts)
		if block == nil {
			break
		}
		// Only use PEM "CERTIFICATE" blocks without extra headers
		if block.Type != CertificateBlockType || len(block.Headers) != 0 {
			continue
		}

		cert, err := x509.ParseCertificate(block.Bytes)
		if err != nil {
			return certs, err
		}

		certs = append(certs, cert)
		ok = true
	}

	if !ok {
		return certs, errors.New("data does not contain any valid RSA or ECDSA certificates")
	}
	return certs, nil
}

// ParsePrivateKeyPEM returns a private key parsed from a PEM block in the supplied data.
// Recognizes PEM blocks for "EC PRIVATE KEY", "RSA PRIVATE KEY", or "PRIVATE KEY"
func ParsePrivateKeyPEM(keyData []byte) (interface{}, error) {
	var privateKeyPemBlock *pem.Block
	for {
		privateKeyPemBlock, keyData = pem.Decode(keyData)
		if privateKeyPemBlock == nil {
			break
		}

		switch privateKeyPemBlock.Type {
		case ECPrivateKeyBlockType:
			// ECDSA Private Key in ASN.1 format
			if key, err := x509.ParseECPrivateKey(privateKeyPemBlock.Bytes); err == nil {
				return key, nil
			}
		case RSAPrivateKeyBlockType:
			// RSA Private Key in PKCS#1 format
			if key, err := x509.ParsePKCS1PrivateKey(privateKeyPemBlock.Bytes); err == nil {
				return key, nil
			}
		case PrivateKeyBlockType:
			// RSA or ECDSA Private Key in unencrypted PKCS#8 format
			if key, err := x509.ParsePKCS8PrivateKey(privateKeyPemBlock.Bytes); err == nil {
				return key, nil
			}
		}

		// tolerate non-key PEM blocks for compatibility with things like "EC PARAMETERS" blocks
		// originally, only the first PEM block was parsed and expected to be a key block
	}

	// we read all the PEM blocks and didn't recognize one
	return nil, fmt.Errorf("data does not contain a valid RSA or ECDSA private key")
}

func EncodeCertPEM(cert *x509.Certificate) []byte {
	block := pem.Block{
		Type:  CertificateBlockType,
		Bytes: cert.Raw,
	}
	return pem.EncodeToMemory(&block)
}

// WriteCertAndKey stores certificate and key at the specified location
func WriteCertAndKey(pkiPath string, name string, cert *x509.Certificate, key crypto.Signer) error {
	if err := WriteKey(pkiPath, name, key); err != nil {
		return err
	}

	return WriteCert(pkiPath, name, cert)
}

// WriteCert stores the given certificate at the given location
func WriteCert(pkiPath, name string, cert *x509.Certificate) error {
	if cert == nil {
		return errors.New("certificate cannot be nil when writing to file")
	}

	certificatePath := pathForCert(pkiPath, name)
	if err := certutil.WriteCert(certificatePath, EncodeCertPEM(cert)); err != nil {
		return err
	}

	return nil
}

// WriteKey stores the given key at the given location
func WriteKey(pkiPath, name string, key crypto.Signer) error {
	if key == nil {
		return errors.New("private key cannot be nil when writing to file")
	}

	privateKeyPath := pathForKey(pkiPath, name)
	encoded, err := keyutil.MarshalPrivateKeyToPEM(key)
	if err != nil {
		return err
	}
	if err := keyutil.WriteKey(privateKeyPath, encoded); err != nil {
		return err
	}

	return nil
}

func pathForCert(pkiPath, name string) string {
	return filepath.Join(pkiPath, fmt.Sprintf("%s.crt", name))
}

func pathForKey(pkiPath, name string) string {
	return filepath.Join(pkiPath, fmt.Sprintf("%s.key", name))
}

// RemoveDuplicateAltNames removes duplicate items in altNames.
func RemoveDuplicateAltNames(altNames *certutil.AltNames) {
	if altNames == nil {
		return
	}

	if altNames.DNSNames != nil {
		altNames.DNSNames = sets.NewString(altNames.DNSNames...).List()
	}

	ipsKeys := make(map[string]struct{})
	var ips []net.IP
	for _, one := range altNames.IPs {
		if _, ok := ipsKeys[one.String()]; !ok {
			ipsKeys[one.String()] = struct{}{}
			ips = append(ips, one)
		}
	}
	altNames.IPs = ips
}
