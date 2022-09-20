// This file is auto-generated. DO NOT EDIT

package jwk

import (
	"bytes"
	"context"
	"crypto/rsa"
	"crypto/x509"
	"fmt"
	"sort"
	"sync"

	"github.com/lestrrat-go/iter/mapiter"
	"github.com/lestrrat-go/jwx/internal/base64"
	"github.com/lestrrat-go/jwx/internal/iter"
	"github.com/lestrrat-go/jwx/internal/json"
	"github.com/lestrrat-go/jwx/internal/pool"
	"github.com/lestrrat-go/jwx/jwa"
	"github.com/pkg/errors"
)

const (
	RSADKey  = "d"
	RSADPKey = "dp"
	RSADQKey = "dq"
	RSAEKey  = "e"
	RSANKey  = "n"
	RSAPKey  = "p"
	RSAQKey  = "q"
	RSAQIKey = "qi"
)

type RSAPrivateKey interface {
	Key
	FromRaw(*rsa.PrivateKey) error
	D() []byte
	DP() []byte
	DQ() []byte
	E() []byte
	N() []byte
	P() []byte
	Q() []byte
	QI() []byte
}

type rsaPrivateKey struct {
	algorithm              *string // https://tools.ietf.org/html/rfc7517#section-4.4
	d                      []byte
	dp                     []byte
	dq                     []byte
	e                      []byte
	keyID                  *string           // https://tools.ietf.org/html/rfc7515#section-4.1.4
	keyUsage               *string           // https://tools.ietf.org/html/rfc7517#section-4.2
	keyops                 *KeyOperationList // https://tools.ietf.org/html/rfc7517#section-4.3
	n                      []byte
	p                      []byte
	q                      []byte
	qi                     []byte
	x509CertChain          *CertificateChain // https://tools.ietf.org/html/rfc7515#section-4.1.6
	x509CertThumbprint     *string           // https://tools.ietf.org/html/rfc7515#section-4.1.7
	x509CertThumbprintS256 *string           // https://tools.ietf.org/html/rfc7515#section-4.1.8
	x509URL                *string           // https://tools.ietf.org/html/rfc7515#section-4.1.5
	privateParams          map[string]interface{}
	mu                     *sync.RWMutex
	dc                     DecodeCtx
}

func NewRSAPrivateKey() RSAPrivateKey {
	return newRSAPrivateKey()
}

func newRSAPrivateKey() *rsaPrivateKey {
	return &rsaPrivateKey{
		mu:            &sync.RWMutex{},
		privateParams: make(map[string]interface{}),
	}
}

func (h rsaPrivateKey) KeyType() jwa.KeyType {
	return jwa.RSA
}

func (h *rsaPrivateKey) Algorithm() string {
	if h.algorithm != nil {
		return *(h.algorithm)
	}
	return ""
}

func (h *rsaPrivateKey) D() []byte {
	return h.d
}

func (h *rsaPrivateKey) DP() []byte {
	return h.dp
}

func (h *rsaPrivateKey) DQ() []byte {
	return h.dq
}

func (h *rsaPrivateKey) E() []byte {
	return h.e
}

func (h *rsaPrivateKey) KeyID() string {
	if h.keyID != nil {
		return *(h.keyID)
	}
	return ""
}

func (h *rsaPrivateKey) KeyUsage() string {
	if h.keyUsage != nil {
		return *(h.keyUsage)
	}
	return ""
}

func (h *rsaPrivateKey) KeyOps() KeyOperationList {
	if h.keyops != nil {
		return *(h.keyops)
	}
	return nil
}

func (h *rsaPrivateKey) N() []byte {
	return h.n
}

func (h *rsaPrivateKey) P() []byte {
	return h.p
}

func (h *rsaPrivateKey) Q() []byte {
	return h.q
}

func (h *rsaPrivateKey) QI() []byte {
	return h.qi
}

func (h *rsaPrivateKey) X509CertChain() []*x509.Certificate {
	if h.x509CertChain != nil {
		return h.x509CertChain.Get()
	}
	return nil
}

func (h *rsaPrivateKey) X509CertThumbprint() string {
	if h.x509CertThumbprint != nil {
		return *(h.x509CertThumbprint)
	}
	return ""
}

func (h *rsaPrivateKey) X509CertThumbprintS256() string {
	if h.x509CertThumbprintS256 != nil {
		return *(h.x509CertThumbprintS256)
	}
	return ""
}

func (h *rsaPrivateKey) X509URL() string {
	if h.x509URL != nil {
		return *(h.x509URL)
	}
	return ""
}

func (h *rsaPrivateKey) makePairs() []*HeaderPair {
	h.mu.RLock()
	defer h.mu.RUnlock()

	var pairs []*HeaderPair
	pairs = append(pairs, &HeaderPair{Key: "kty", Value: jwa.RSA})
	if h.algorithm != nil {
		pairs = append(pairs, &HeaderPair{Key: AlgorithmKey, Value: *(h.algorithm)})
	}
	if h.d != nil {
		pairs = append(pairs, &HeaderPair{Key: RSADKey, Value: h.d})
	}
	if h.dp != nil {
		pairs = append(pairs, &HeaderPair{Key: RSADPKey, Value: h.dp})
	}
	if h.dq != nil {
		pairs = append(pairs, &HeaderPair{Key: RSADQKey, Value: h.dq})
	}
	if h.e != nil {
		pairs = append(pairs, &HeaderPair{Key: RSAEKey, Value: h.e})
	}
	if h.keyID != nil {
		pairs = append(pairs, &HeaderPair{Key: KeyIDKey, Value: *(h.keyID)})
	}
	if h.keyUsage != nil {
		pairs = append(pairs, &HeaderPair{Key: KeyUsageKey, Value: *(h.keyUsage)})
	}
	if h.keyops != nil {
		pairs = append(pairs, &HeaderPair{Key: KeyOpsKey, Value: *(h.keyops)})
	}
	if h.n != nil {
		pairs = append(pairs, &HeaderPair{Key: RSANKey, Value: h.n})
	}
	if h.p != nil {
		pairs = append(pairs, &HeaderPair{Key: RSAPKey, Value: h.p})
	}
	if h.q != nil {
		pairs = append(pairs, &HeaderPair{Key: RSAQKey, Value: h.q})
	}
	if h.qi != nil {
		pairs = append(pairs, &HeaderPair{Key: RSAQIKey, Value: h.qi})
	}
	if h.x509CertChain != nil {
		pairs = append(pairs, &HeaderPair{Key: X509CertChainKey, Value: *(h.x509CertChain)})
	}
	if h.x509CertThumbprint != nil {
		pairs = append(pairs, &HeaderPair{Key: X509CertThumbprintKey, Value: *(h.x509CertThumbprint)})
	}
	if h.x509CertThumbprintS256 != nil {
		pairs = append(pairs, &HeaderPair{Key: X509CertThumbprintS256Key, Value: *(h.x509CertThumbprintS256)})
	}
	if h.x509URL != nil {
		pairs = append(pairs, &HeaderPair{Key: X509URLKey, Value: *(h.x509URL)})
	}
	for k, v := range h.privateParams {
		pairs = append(pairs, &HeaderPair{Key: k, Value: v})
	}
	return pairs
}

func (h *rsaPrivateKey) PrivateParams() map[string]interface{} {
	return h.privateParams
}

func (h *rsaPrivateKey) Get(name string) (interface{}, bool) {
	h.mu.RLock()
	defer h.mu.RUnlock()
	switch name {
	case KeyTypeKey:
		return h.KeyType(), true
	case AlgorithmKey:
		if h.algorithm == nil {
			return nil, false
		}
		return *(h.algorithm), true
	case RSADKey:
		if h.d == nil {
			return nil, false
		}
		return h.d, true
	case RSADPKey:
		if h.dp == nil {
			return nil, false
		}
		return h.dp, true
	case RSADQKey:
		if h.dq == nil {
			return nil, false
		}
		return h.dq, true
	case RSAEKey:
		if h.e == nil {
			return nil, false
		}
		return h.e, true
	case KeyIDKey:
		if h.keyID == nil {
			return nil, false
		}
		return *(h.keyID), true
	case KeyUsageKey:
		if h.keyUsage == nil {
			return nil, false
		}
		return *(h.keyUsage), true
	case KeyOpsKey:
		if h.keyops == nil {
			return nil, false
		}
		return *(h.keyops), true
	case RSANKey:
		if h.n == nil {
			return nil, false
		}
		return h.n, true
	case RSAPKey:
		if h.p == nil {
			return nil, false
		}
		return h.p, true
	case RSAQKey:
		if h.q == nil {
			return nil, false
		}
		return h.q, true
	case RSAQIKey:
		if h.qi == nil {
			return nil, false
		}
		return h.qi, true
	case X509CertChainKey:
		if h.x509CertChain == nil {
			return nil, false
		}
		return h.x509CertChain.Get(), true
	case X509CertThumbprintKey:
		if h.x509CertThumbprint == nil {
			return nil, false
		}
		return *(h.x509CertThumbprint), true
	case X509CertThumbprintS256Key:
		if h.x509CertThumbprintS256 == nil {
			return nil, false
		}
		return *(h.x509CertThumbprintS256), true
	case X509URLKey:
		if h.x509URL == nil {
			return nil, false
		}
		return *(h.x509URL), true
	default:
		v, ok := h.privateParams[name]
		return v, ok
	}
}

func (h *rsaPrivateKey) Set(name string, value interface{}) error {
	h.mu.Lock()
	defer h.mu.Unlock()
	return h.setNoLock(name, value)
}

func (h *rsaPrivateKey) setNoLock(name string, value interface{}) error {
	switch name {
	case "kty":
		return nil
	case AlgorithmKey:
		switch v := value.(type) {
		case string:
			h.algorithm = &v
		case fmt.Stringer:
			tmp := v.String()
			h.algorithm = &tmp
		default:
			return errors.Errorf(`invalid type for %s key: %T`, AlgorithmKey, value)
		}
		return nil
	case RSADKey:
		if v, ok := value.([]byte); ok {
			h.d = v
			return nil
		}
		return errors.Errorf(`invalid value for %s key: %T`, RSADKey, value)
	case RSADPKey:
		if v, ok := value.([]byte); ok {
			h.dp = v
			return nil
		}
		return errors.Errorf(`invalid value for %s key: %T`, RSADPKey, value)
	case RSADQKey:
		if v, ok := value.([]byte); ok {
			h.dq = v
			return nil
		}
		return errors.Errorf(`invalid value for %s key: %T`, RSADQKey, value)
	case RSAEKey:
		if v, ok := value.([]byte); ok {
			h.e = v
			return nil
		}
		return errors.Errorf(`invalid value for %s key: %T`, RSAEKey, value)
	case KeyIDKey:
		if v, ok := value.(string); ok {
			h.keyID = &v
			return nil
		}
		return errors.Errorf(`invalid value for %s key: %T`, KeyIDKey, value)
	case KeyUsageKey:
		switch v := value.(type) {
		case KeyUsageType:
			switch v {
			case ForSignature, ForEncryption:
				tmp := v.String()
				h.keyUsage = &tmp
			default:
				return errors.Errorf(`invalid key usage type %s`, v)
			}
		case string:
			h.keyUsage = &v
		default:
			return errors.Errorf(`invalid key usage type %s`, v)
		}
	case KeyOpsKey:
		var acceptor KeyOperationList
		if err := acceptor.Accept(value); err != nil {
			return errors.Wrapf(err, `invalid value for %s key`, KeyOpsKey)
		}
		h.keyops = &acceptor
		return nil
	case RSANKey:
		if v, ok := value.([]byte); ok {
			h.n = v
			return nil
		}
		return errors.Errorf(`invalid value for %s key: %T`, RSANKey, value)
	case RSAPKey:
		if v, ok := value.([]byte); ok {
			h.p = v
			return nil
		}
		return errors.Errorf(`invalid value for %s key: %T`, RSAPKey, value)
	case RSAQKey:
		if v, ok := value.([]byte); ok {
			h.q = v
			return nil
		}
		return errors.Errorf(`invalid value for %s key: %T`, RSAQKey, value)
	case RSAQIKey:
		if v, ok := value.([]byte); ok {
			h.qi = v
			return nil
		}
		return errors.Errorf(`invalid value for %s key: %T`, RSAQIKey, value)
	case X509CertChainKey:
		var acceptor CertificateChain
		if err := acceptor.Accept(value); err != nil {
			return errors.Wrapf(err, `invalid value for %s key`, X509CertChainKey)
		}
		h.x509CertChain = &acceptor
		return nil
	case X509CertThumbprintKey:
		if v, ok := value.(string); ok {
			h.x509CertThumbprint = &v
			return nil
		}
		return errors.Errorf(`invalid value for %s key: %T`, X509CertThumbprintKey, value)
	case X509CertThumbprintS256Key:
		if v, ok := value.(string); ok {
			h.x509CertThumbprintS256 = &v
			return nil
		}
		return errors.Errorf(`invalid value for %s key: %T`, X509CertThumbprintS256Key, value)
	case X509URLKey:
		if v, ok := value.(string); ok {
			h.x509URL = &v
			return nil
		}
		return errors.Errorf(`invalid value for %s key: %T`, X509URLKey, value)
	default:
		if h.privateParams == nil {
			h.privateParams = map[string]interface{}{}
		}
		h.privateParams[name] = value
	}
	return nil
}

func (k *rsaPrivateKey) Remove(key string) error {
	k.mu.Lock()
	defer k.mu.Unlock()
	switch key {
	case AlgorithmKey:
		k.algorithm = nil
	case RSADKey:
		k.d = nil
	case RSADPKey:
		k.dp = nil
	case RSADQKey:
		k.dq = nil
	case RSAEKey:
		k.e = nil
	case KeyIDKey:
		k.keyID = nil
	case KeyUsageKey:
		k.keyUsage = nil
	case KeyOpsKey:
		k.keyops = nil
	case RSANKey:
		k.n = nil
	case RSAPKey:
		k.p = nil
	case RSAQKey:
		k.q = nil
	case RSAQIKey:
		k.qi = nil
	case X509CertChainKey:
		k.x509CertChain = nil
	case X509CertThumbprintKey:
		k.x509CertThumbprint = nil
	case X509CertThumbprintS256Key:
		k.x509CertThumbprintS256 = nil
	case X509URLKey:
		k.x509URL = nil
	default:
		delete(k.privateParams, key)
	}
	return nil
}

func (k *rsaPrivateKey) Clone() (Key, error) {
	return cloneKey(k)
}

func (k *rsaPrivateKey) DecodeCtx() DecodeCtx {
	k.mu.RLock()
	defer k.mu.RUnlock()
	return k.dc
}

func (k *rsaPrivateKey) SetDecodeCtx(dc DecodeCtx) {
	k.mu.Lock()
	defer k.mu.Unlock()
	k.dc = dc
}

func (h *rsaPrivateKey) UnmarshalJSON(buf []byte) error {
	h.algorithm = nil
	h.d = nil
	h.dp = nil
	h.dq = nil
	h.e = nil
	h.keyID = nil
	h.keyUsage = nil
	h.keyops = nil
	h.n = nil
	h.p = nil
	h.q = nil
	h.qi = nil
	h.x509CertChain = nil
	h.x509CertThumbprint = nil
	h.x509CertThumbprintS256 = nil
	h.x509URL = nil
	dec := json.NewDecoder(bytes.NewReader(buf))
LOOP:
	for {
		tok, err := dec.Token()
		if err != nil {
			return errors.Wrap(err, `error reading token`)
		}
		switch tok := tok.(type) {
		case json.Delim:
			// Assuming we're doing everything correctly, we should ONLY
			// get either '{' or '}' here.
			if tok == '}' { // End of object
				break LOOP
			} else if tok != '{' {
				return errors.Errorf(`expected '{', but got '%c'`, tok)
			}
		case string: // Objects can only have string keys
			switch tok {
			case KeyTypeKey:
				val, err := json.ReadNextStringToken(dec)
				if err != nil {
					return errors.Wrap(err, `error reading token`)
				}
				if val != jwa.RSA.String() {
					return errors.Errorf(`invalid kty value for RSAPublicKey (%s)`, val)
				}
			case AlgorithmKey:
				if err := json.AssignNextStringToken(&h.algorithm, dec); err != nil {
					return errors.Wrapf(err, `failed to decode value for key %s`, AlgorithmKey)
				}
			case RSADKey:
				if err := json.AssignNextBytesToken(&h.d, dec); err != nil {
					return errors.Wrapf(err, `failed to decode value for key %s`, RSADKey)
				}
			case RSADPKey:
				if err := json.AssignNextBytesToken(&h.dp, dec); err != nil {
					return errors.Wrapf(err, `failed to decode value for key %s`, RSADPKey)
				}
			case RSADQKey:
				if err := json.AssignNextBytesToken(&h.dq, dec); err != nil {
					return errors.Wrapf(err, `failed to decode value for key %s`, RSADQKey)
				}
			case RSAEKey:
				if err := json.AssignNextBytesToken(&h.e, dec); err != nil {
					return errors.Wrapf(err, `failed to decode value for key %s`, RSAEKey)
				}
			case KeyIDKey:
				if err := json.AssignNextStringToken(&h.keyID, dec); err != nil {
					return errors.Wrapf(err, `failed to decode value for key %s`, KeyIDKey)
				}
			case KeyUsageKey:
				if err := json.AssignNextStringToken(&h.keyUsage, dec); err != nil {
					return errors.Wrapf(err, `failed to decode value for key %s`, KeyUsageKey)
				}
			case KeyOpsKey:
				var decoded KeyOperationList
				if err := dec.Decode(&decoded); err != nil {
					return errors.Wrapf(err, `failed to decode value for key %s`, KeyOpsKey)
				}
				h.keyops = &decoded
			case RSANKey:
				if err := json.AssignNextBytesToken(&h.n, dec); err != nil {
					return errors.Wrapf(err, `failed to decode value for key %s`, RSANKey)
				}
			case RSAPKey:
				if err := json.AssignNextBytesToken(&h.p, dec); err != nil {
					return errors.Wrapf(err, `failed to decode value for key %s`, RSAPKey)
				}
			case RSAQKey:
				if err := json.AssignNextBytesToken(&h.q, dec); err != nil {
					return errors.Wrapf(err, `failed to decode value for key %s`, RSAQKey)
				}
			case RSAQIKey:
				if err := json.AssignNextBytesToken(&h.qi, dec); err != nil {
					return errors.Wrapf(err, `failed to decode value for key %s`, RSAQIKey)
				}
			case X509CertChainKey:
				var decoded CertificateChain
				if err := dec.Decode(&decoded); err != nil {
					return errors.Wrapf(err, `failed to decode value for key %s`, X509CertChainKey)
				}
				h.x509CertChain = &decoded
			case X509CertThumbprintKey:
				if err := json.AssignNextStringToken(&h.x509CertThumbprint, dec); err != nil {
					return errors.Wrapf(err, `failed to decode value for key %s`, X509CertThumbprintKey)
				}
			case X509CertThumbprintS256Key:
				if err := json.AssignNextStringToken(&h.x509CertThumbprintS256, dec); err != nil {
					return errors.Wrapf(err, `failed to decode value for key %s`, X509CertThumbprintS256Key)
				}
			case X509URLKey:
				if err := json.AssignNextStringToken(&h.x509URL, dec); err != nil {
					return errors.Wrapf(err, `failed to decode value for key %s`, X509URLKey)
				}
			default:
				if dc := h.dc; dc != nil {
					if localReg := dc.Registry(); localReg != nil {
						decoded, err := localReg.Decode(dec, tok)
						if err == nil {
							h.setNoLock(tok, decoded)
							continue
						}
					}
				}
				decoded, err := registry.Decode(dec, tok)
				if err == nil {
					h.setNoLock(tok, decoded)
					continue
				}
				return errors.Wrapf(err, `could not decode field %s`, tok)
			}
		default:
			return errors.Errorf(`invalid token %T`, tok)
		}
	}
	if h.d == nil {
		return errors.Errorf(`required field d is missing`)
	}
	if h.e == nil {
		return errors.Errorf(`required field e is missing`)
	}
	if h.n == nil {
		return errors.Errorf(`required field n is missing`)
	}
	if h.p == nil {
		return errors.Errorf(`required field p is missing`)
	}
	if h.q == nil {
		return errors.Errorf(`required field q is missing`)
	}
	return nil
}

func (h rsaPrivateKey) MarshalJSON() ([]byte, error) {
	data := make(map[string]interface{})
	fields := make([]string, 0, 16)
	for _, pair := range h.makePairs() {
		fields = append(fields, pair.Key.(string))
		data[pair.Key.(string)] = pair.Value
	}

	sort.Strings(fields)
	buf := pool.GetBytesBuffer()
	defer pool.ReleaseBytesBuffer(buf)
	buf.WriteByte('{')
	enc := json.NewEncoder(buf)
	for i, f := range fields {
		if i > 0 {
			buf.WriteRune(',')
		}
		buf.WriteRune('"')
		buf.WriteString(f)
		buf.WriteString(`":`)
		v := data[f]
		switch v := v.(type) {
		case []byte:
			buf.WriteRune('"')
			buf.WriteString(base64.EncodeToString(v))
			buf.WriteRune('"')
		default:
			if err := enc.Encode(v); err != nil {
				return nil, errors.Wrapf(err, `failed to encode value for field %s`, f)
			}
			buf.Truncate(buf.Len() - 1)
		}
	}
	buf.WriteByte('}')
	ret := make([]byte, buf.Len())
	copy(ret, buf.Bytes())
	return ret, nil
}

func (h *rsaPrivateKey) Iterate(ctx context.Context) HeaderIterator {
	pairs := h.makePairs()
	ch := make(chan *HeaderPair, len(pairs))
	go func(ctx context.Context, ch chan *HeaderPair, pairs []*HeaderPair) {
		defer close(ch)
		for _, pair := range pairs {
			select {
			case <-ctx.Done():
				return
			case ch <- pair:
			}
		}
	}(ctx, ch, pairs)
	return mapiter.New(ch)
}

func (h *rsaPrivateKey) Walk(ctx context.Context, visitor HeaderVisitor) error {
	return iter.WalkMap(ctx, h, visitor)
}

func (h *rsaPrivateKey) AsMap(ctx context.Context) (map[string]interface{}, error) {
	return iter.AsMap(ctx, h)
}

type RSAPublicKey interface {
	Key
	FromRaw(*rsa.PublicKey) error
	E() []byte
	N() []byte
}

type rsaPublicKey struct {
	algorithm              *string // https://tools.ietf.org/html/rfc7517#section-4.4
	e                      []byte
	keyID                  *string           // https://tools.ietf.org/html/rfc7515#section-4.1.4
	keyUsage               *string           // https://tools.ietf.org/html/rfc7517#section-4.2
	keyops                 *KeyOperationList // https://tools.ietf.org/html/rfc7517#section-4.3
	n                      []byte
	x509CertChain          *CertificateChain // https://tools.ietf.org/html/rfc7515#section-4.1.6
	x509CertThumbprint     *string           // https://tools.ietf.org/html/rfc7515#section-4.1.7
	x509CertThumbprintS256 *string           // https://tools.ietf.org/html/rfc7515#section-4.1.8
	x509URL                *string           // https://tools.ietf.org/html/rfc7515#section-4.1.5
	privateParams          map[string]interface{}
	mu                     *sync.RWMutex
	dc                     DecodeCtx
}

func NewRSAPublicKey() RSAPublicKey {
	return newRSAPublicKey()
}

func newRSAPublicKey() *rsaPublicKey {
	return &rsaPublicKey{
		mu:            &sync.RWMutex{},
		privateParams: make(map[string]interface{}),
	}
}

func (h rsaPublicKey) KeyType() jwa.KeyType {
	return jwa.RSA
}

func (h *rsaPublicKey) Algorithm() string {
	if h.algorithm != nil {
		return *(h.algorithm)
	}
	return ""
}

func (h *rsaPublicKey) E() []byte {
	return h.e
}

func (h *rsaPublicKey) KeyID() string {
	if h.keyID != nil {
		return *(h.keyID)
	}
	return ""
}

func (h *rsaPublicKey) KeyUsage() string {
	if h.keyUsage != nil {
		return *(h.keyUsage)
	}
	return ""
}

func (h *rsaPublicKey) KeyOps() KeyOperationList {
	if h.keyops != nil {
		return *(h.keyops)
	}
	return nil
}

func (h *rsaPublicKey) N() []byte {
	return h.n
}

func (h *rsaPublicKey) X509CertChain() []*x509.Certificate {
	if h.x509CertChain != nil {
		return h.x509CertChain.Get()
	}
	return nil
}

func (h *rsaPublicKey) X509CertThumbprint() string {
	if h.x509CertThumbprint != nil {
		return *(h.x509CertThumbprint)
	}
	return ""
}

func (h *rsaPublicKey) X509CertThumbprintS256() string {
	if h.x509CertThumbprintS256 != nil {
		return *(h.x509CertThumbprintS256)
	}
	return ""
}

func (h *rsaPublicKey) X509URL() string {
	if h.x509URL != nil {
		return *(h.x509URL)
	}
	return ""
}

func (h *rsaPublicKey) makePairs() []*HeaderPair {
	h.mu.RLock()
	defer h.mu.RUnlock()

	var pairs []*HeaderPair
	pairs = append(pairs, &HeaderPair{Key: "kty", Value: jwa.RSA})
	if h.algorithm != nil {
		pairs = append(pairs, &HeaderPair{Key: AlgorithmKey, Value: *(h.algorithm)})
	}
	if h.e != nil {
		pairs = append(pairs, &HeaderPair{Key: RSAEKey, Value: h.e})
	}
	if h.keyID != nil {
		pairs = append(pairs, &HeaderPair{Key: KeyIDKey, Value: *(h.keyID)})
	}
	if h.keyUsage != nil {
		pairs = append(pairs, &HeaderPair{Key: KeyUsageKey, Value: *(h.keyUsage)})
	}
	if h.keyops != nil {
		pairs = append(pairs, &HeaderPair{Key: KeyOpsKey, Value: *(h.keyops)})
	}
	if h.n != nil {
		pairs = append(pairs, &HeaderPair{Key: RSANKey, Value: h.n})
	}
	if h.x509CertChain != nil {
		pairs = append(pairs, &HeaderPair{Key: X509CertChainKey, Value: *(h.x509CertChain)})
	}
	if h.x509CertThumbprint != nil {
		pairs = append(pairs, &HeaderPair{Key: X509CertThumbprintKey, Value: *(h.x509CertThumbprint)})
	}
	if h.x509CertThumbprintS256 != nil {
		pairs = append(pairs, &HeaderPair{Key: X509CertThumbprintS256Key, Value: *(h.x509CertThumbprintS256)})
	}
	if h.x509URL != nil {
		pairs = append(pairs, &HeaderPair{Key: X509URLKey, Value: *(h.x509URL)})
	}
	for k, v := range h.privateParams {
		pairs = append(pairs, &HeaderPair{Key: k, Value: v})
	}
	return pairs
}

func (h *rsaPublicKey) PrivateParams() map[string]interface{} {
	return h.privateParams
}

func (h *rsaPublicKey) Get(name string) (interface{}, bool) {
	h.mu.RLock()
	defer h.mu.RUnlock()
	switch name {
	case KeyTypeKey:
		return h.KeyType(), true
	case AlgorithmKey:
		if h.algorithm == nil {
			return nil, false
		}
		return *(h.algorithm), true
	case RSAEKey:
		if h.e == nil {
			return nil, false
		}
		return h.e, true
	case KeyIDKey:
		if h.keyID == nil {
			return nil, false
		}
		return *(h.keyID), true
	case KeyUsageKey:
		if h.keyUsage == nil {
			return nil, false
		}
		return *(h.keyUsage), true
	case KeyOpsKey:
		if h.keyops == nil {
			return nil, false
		}
		return *(h.keyops), true
	case RSANKey:
		if h.n == nil {
			return nil, false
		}
		return h.n, true
	case X509CertChainKey:
		if h.x509CertChain == nil {
			return nil, false
		}
		return h.x509CertChain.Get(), true
	case X509CertThumbprintKey:
		if h.x509CertThumbprint == nil {
			return nil, false
		}
		return *(h.x509CertThumbprint), true
	case X509CertThumbprintS256Key:
		if h.x509CertThumbprintS256 == nil {
			return nil, false
		}
		return *(h.x509CertThumbprintS256), true
	case X509URLKey:
		if h.x509URL == nil {
			return nil, false
		}
		return *(h.x509URL), true
	default:
		v, ok := h.privateParams[name]
		return v, ok
	}
}

func (h *rsaPublicKey) Set(name string, value interface{}) error {
	h.mu.Lock()
	defer h.mu.Unlock()
	return h.setNoLock(name, value)
}

func (h *rsaPublicKey) setNoLock(name string, value interface{}) error {
	switch name {
	case "kty":
		return nil
	case AlgorithmKey:
		switch v := value.(type) {
		case string:
			h.algorithm = &v
		case fmt.Stringer:
			tmp := v.String()
			h.algorithm = &tmp
		default:
			return errors.Errorf(`invalid type for %s key: %T`, AlgorithmKey, value)
		}
		return nil
	case RSAEKey:
		if v, ok := value.([]byte); ok {
			h.e = v
			return nil
		}
		return errors.Errorf(`invalid value for %s key: %T`, RSAEKey, value)
	case KeyIDKey:
		if v, ok := value.(string); ok {
			h.keyID = &v
			return nil
		}
		return errors.Errorf(`invalid value for %s key: %T`, KeyIDKey, value)
	case KeyUsageKey:
		switch v := value.(type) {
		case KeyUsageType:
			switch v {
			case ForSignature, ForEncryption:
				tmp := v.String()
				h.keyUsage = &tmp
			default:
				return errors.Errorf(`invalid key usage type %s`, v)
			}
		case string:
			h.keyUsage = &v
		default:
			return errors.Errorf(`invalid key usage type %s`, v)
		}
	case KeyOpsKey:
		var acceptor KeyOperationList
		if err := acceptor.Accept(value); err != nil {
			return errors.Wrapf(err, `invalid value for %s key`, KeyOpsKey)
		}
		h.keyops = &acceptor
		return nil
	case RSANKey:
		if v, ok := value.([]byte); ok {
			h.n = v
			return nil
		}
		return errors.Errorf(`invalid value for %s key: %T`, RSANKey, value)
	case X509CertChainKey:
		var acceptor CertificateChain
		if err := acceptor.Accept(value); err != nil {
			return errors.Wrapf(err, `invalid value for %s key`, X509CertChainKey)
		}
		h.x509CertChain = &acceptor
		return nil
	case X509CertThumbprintKey:
		if v, ok := value.(string); ok {
			h.x509CertThumbprint = &v
			return nil
		}
		return errors.Errorf(`invalid value for %s key: %T`, X509CertThumbprintKey, value)
	case X509CertThumbprintS256Key:
		if v, ok := value.(string); ok {
			h.x509CertThumbprintS256 = &v
			return nil
		}
		return errors.Errorf(`invalid value for %s key: %T`, X509CertThumbprintS256Key, value)
	case X509URLKey:
		if v, ok := value.(string); ok {
			h.x509URL = &v
			return nil
		}
		return errors.Errorf(`invalid value for %s key: %T`, X509URLKey, value)
	default:
		if h.privateParams == nil {
			h.privateParams = map[string]interface{}{}
		}
		h.privateParams[name] = value
	}
	return nil
}

func (k *rsaPublicKey) Remove(key string) error {
	k.mu.Lock()
	defer k.mu.Unlock()
	switch key {
	case AlgorithmKey:
		k.algorithm = nil
	case RSAEKey:
		k.e = nil
	case KeyIDKey:
		k.keyID = nil
	case KeyUsageKey:
		k.keyUsage = nil
	case KeyOpsKey:
		k.keyops = nil
	case RSANKey:
		k.n = nil
	case X509CertChainKey:
		k.x509CertChain = nil
	case X509CertThumbprintKey:
		k.x509CertThumbprint = nil
	case X509CertThumbprintS256Key:
		k.x509CertThumbprintS256 = nil
	case X509URLKey:
		k.x509URL = nil
	default:
		delete(k.privateParams, key)
	}
	return nil
}

func (k *rsaPublicKey) Clone() (Key, error) {
	return cloneKey(k)
}

func (k *rsaPublicKey) DecodeCtx() DecodeCtx {
	k.mu.RLock()
	defer k.mu.RUnlock()
	return k.dc
}

func (k *rsaPublicKey) SetDecodeCtx(dc DecodeCtx) {
	k.mu.Lock()
	defer k.mu.Unlock()
	k.dc = dc
}

func (h *rsaPublicKey) UnmarshalJSON(buf []byte) error {
	h.algorithm = nil
	h.e = nil
	h.keyID = nil
	h.keyUsage = nil
	h.keyops = nil
	h.n = nil
	h.x509CertChain = nil
	h.x509CertThumbprint = nil
	h.x509CertThumbprintS256 = nil
	h.x509URL = nil
	dec := json.NewDecoder(bytes.NewReader(buf))
LOOP:
	for {
		tok, err := dec.Token()
		if err != nil {
			return errors.Wrap(err, `error reading token`)
		}
		switch tok := tok.(type) {
		case json.Delim:
			// Assuming we're doing everything correctly, we should ONLY
			// get either '{' or '}' here.
			if tok == '}' { // End of object
				break LOOP
			} else if tok != '{' {
				return errors.Errorf(`expected '{', but got '%c'`, tok)
			}
		case string: // Objects can only have string keys
			switch tok {
			case KeyTypeKey:
				val, err := json.ReadNextStringToken(dec)
				if err != nil {
					return errors.Wrap(err, `error reading token`)
				}
				if val != jwa.RSA.String() {
					return errors.Errorf(`invalid kty value for RSAPublicKey (%s)`, val)
				}
			case AlgorithmKey:
				if err := json.AssignNextStringToken(&h.algorithm, dec); err != nil {
					return errors.Wrapf(err, `failed to decode value for key %s`, AlgorithmKey)
				}
			case RSAEKey:
				if err := json.AssignNextBytesToken(&h.e, dec); err != nil {
					return errors.Wrapf(err, `failed to decode value for key %s`, RSAEKey)
				}
			case KeyIDKey:
				if err := json.AssignNextStringToken(&h.keyID, dec); err != nil {
					return errors.Wrapf(err, `failed to decode value for key %s`, KeyIDKey)
				}
			case KeyUsageKey:
				if err := json.AssignNextStringToken(&h.keyUsage, dec); err != nil {
					return errors.Wrapf(err, `failed to decode value for key %s`, KeyUsageKey)
				}
			case KeyOpsKey:
				var decoded KeyOperationList
				if err := dec.Decode(&decoded); err != nil {
					return errors.Wrapf(err, `failed to decode value for key %s`, KeyOpsKey)
				}
				h.keyops = &decoded
			case RSANKey:
				if err := json.AssignNextBytesToken(&h.n, dec); err != nil {
					return errors.Wrapf(err, `failed to decode value for key %s`, RSANKey)
				}
			case X509CertChainKey:
				var decoded CertificateChain
				if err := dec.Decode(&decoded); err != nil {
					return errors.Wrapf(err, `failed to decode value for key %s`, X509CertChainKey)
				}
				h.x509CertChain = &decoded
			case X509CertThumbprintKey:
				if err := json.AssignNextStringToken(&h.x509CertThumbprint, dec); err != nil {
					return errors.Wrapf(err, `failed to decode value for key %s`, X509CertThumbprintKey)
				}
			case X509CertThumbprintS256Key:
				if err := json.AssignNextStringToken(&h.x509CertThumbprintS256, dec); err != nil {
					return errors.Wrapf(err, `failed to decode value for key %s`, X509CertThumbprintS256Key)
				}
			case X509URLKey:
				if err := json.AssignNextStringToken(&h.x509URL, dec); err != nil {
					return errors.Wrapf(err, `failed to decode value for key %s`, X509URLKey)
				}
			default:
				if dc := h.dc; dc != nil {
					if localReg := dc.Registry(); localReg != nil {
						decoded, err := localReg.Decode(dec, tok)
						if err == nil {
							h.setNoLock(tok, decoded)
							continue
						}
					}
				}
				decoded, err := registry.Decode(dec, tok)
				if err == nil {
					h.setNoLock(tok, decoded)
					continue
				}
				return errors.Wrapf(err, `could not decode field %s`, tok)
			}
		default:
			return errors.Errorf(`invalid token %T`, tok)
		}
	}
	if h.e == nil {
		return errors.Errorf(`required field e is missing`)
	}
	if h.n == nil {
		return errors.Errorf(`required field n is missing`)
	}
	return nil
}

func (h rsaPublicKey) MarshalJSON() ([]byte, error) {
	data := make(map[string]interface{})
	fields := make([]string, 0, 10)
	for _, pair := range h.makePairs() {
		fields = append(fields, pair.Key.(string))
		data[pair.Key.(string)] = pair.Value
	}

	sort.Strings(fields)
	buf := pool.GetBytesBuffer()
	defer pool.ReleaseBytesBuffer(buf)
	buf.WriteByte('{')
	enc := json.NewEncoder(buf)
	for i, f := range fields {
		if i > 0 {
			buf.WriteRune(',')
		}
		buf.WriteRune('"')
		buf.WriteString(f)
		buf.WriteString(`":`)
		v := data[f]
		switch v := v.(type) {
		case []byte:
			buf.WriteRune('"')
			buf.WriteString(base64.EncodeToString(v))
			buf.WriteRune('"')
		default:
			if err := enc.Encode(v); err != nil {
				return nil, errors.Wrapf(err, `failed to encode value for field %s`, f)
			}
			buf.Truncate(buf.Len() - 1)
		}
	}
	buf.WriteByte('}')
	ret := make([]byte, buf.Len())
	copy(ret, buf.Bytes())
	return ret, nil
}

func (h *rsaPublicKey) Iterate(ctx context.Context) HeaderIterator {
	pairs := h.makePairs()
	ch := make(chan *HeaderPair, len(pairs))
	go func(ctx context.Context, ch chan *HeaderPair, pairs []*HeaderPair) {
		defer close(ch)
		for _, pair := range pairs {
			select {
			case <-ctx.Done():
				return
			case ch <- pair:
			}
		}
	}(ctx, ch, pairs)
	return mapiter.New(ch)
}

func (h *rsaPublicKey) Walk(ctx context.Context, visitor HeaderVisitor) error {
	return iter.WalkMap(ctx, h, visitor)
}

func (h *rsaPublicKey) AsMap(ctx context.Context) (map[string]interface{}, error) {
	return iter.AsMap(ctx, h)
}
