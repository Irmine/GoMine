package types

type ChainDataKeys struct {
	RawChains []string `json:"chain"`
	Chains    []Chain
}

type Chain struct {
	Header    ChainHeader
	Payload   ChainPayload
	Signature string
}

type ChainHeader struct {
	X5u string `json:"x5u"`
	Alg string `json:"alg"`

	Raw string
}

type ChainPayload struct {
	CertificateAuthority bool   `json:"certificateAuthority"`
	ExpirationTime       int64  `json:"exp"`
	IdentityPublicKey    string `json:"identityPublicKey"`
	NotBefore            int64  `json:"nbf"`
	RandomNonce          int    `json:"randomNonce"`
	Issuer               string `json:"iss"`
	IssuedAt             int64  `json:"iat"`

	Raw string
}

type WebTokenKeys struct {
	ExtraData         map[string]interface{} `json:"extraData"`
	IdentityPublicKey string                 `json:"identityPublicKey"`
}

type ClientDataKeys struct {
	ClientRandomId   int    `json:"ClientRandomId"`
	ServerAddress    string `json:"ServerAddress"`
	LanguageCode     string `json:"LanguageCode"`
	SkinId           string `json:"SkinId"`
	SkinData         string `json:"SkinData"`
	CapeData         string `json:"CapeData"`
	GeometryId       string `json:"SkinGeometryName"`
	GeometryData     string `json:"SkinGeometry"`
	CurrentInputMode string `json:"CurrentInputMode"`
	DefaultInputMode string `json:"DefaultInputMode"`
	DeviceModel      string `json:"DeviceModel"`
	DeviceOS         int    `json:"DeviceOS"`
	GameVersion      string `json:"GameVersion"`
	GuiScale         int    `json:"GuiScale"`
	UIProfile        int    `json:"UIProfile"`
	ThirdPartyName   string `json:"ThirdPartyName"`
}
