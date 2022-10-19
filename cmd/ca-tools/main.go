package main

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"crypto/x509/pkix"
	"fmt"
	certutil "k8s.io/client-go/util/cert"
	"math"
	"math/big"
	"net"
	"time"
)

func main() {
	TryLoadCertAndKey()
	ca, key, err := TryLoadCertAndKey()
	if err != nil {
		panic(err)
	}
	pemCert, pemKey, err := GenerateCertAndKey(ca, key)
	err = WriteCertAndKey("/tmp/", "boncloud", pemCert, pemKey)
	if err != nil {
		fmt.Println(err)
	}
}
func GenerateCertAndKey(caCert *x509.Certificate, caKey crypto.Signer) (*x509.Certificate, crypto.Signer, error) {
	pemKey, err := rsa.GenerateKey(rand.Reader, 2048)

	serial, err := rand.Int(rand.Reader, new(big.Int).SetInt64(math.MaxInt64))

	keyUsage := x509.KeyUsageKeyEncipherment | x509.KeyUsageDigitalSignature

	altNames := &certutil.AltNames{
		DNSNames: []string{
			cfg.NodeRegistration.Name,
			"kubernetes",
			"kubernetes.default",
			"kubernetes.default.svc",
			fmt.Sprintf("kubernetes.default.svc.%s", cfg.Networking.DNSDomain),
		},
		IPs: []net.IP{
			internalAPIServerVirtualIP,
			advertiseAddress,
		},
	}
	RemoveDuplicateAltNames(altNames)

	notAfter := time.Now().Add(time.Hour * 24 * 365 * 10).UTC()

	certTmpl := x509.Certificate{
		Subject: pkix.Name{
			CommonName:   cfg.CommonName,
			Organization: cfg.Organization,
		},
		DNSNames:              altNames.DNSNames,
		IPAddresses:           altNames.IPs,
		SerialNumber:          serial,
		NotBefore:             caCert.NotBefore,
		NotAfter:              notAfter,
		KeyUsage:              keyUsage,
		ExtKeyUsage:           []x509.ExtKeyUsage{x509.ExtKeyUsageAny},
		BasicConstraintsValid: true,
		IsCA:                  false,
	}

	certDERBytes, err := x509.CreateCertificate(rand.Reader, &certTmpl, caCert, pemKey.Public(), caKey)
	if err != nil {
		panic(err)
	}
	pemCert, err := x509.ParseCertificate(certDERBytes)
	return pemCert, pemKey, err
}
