package permission

import (
	"strings"

	"github.com/kyleu/rituals.dev/app/auth"
	"github.com/kyleu/rituals.dev/app/util"
)

func (s *Service) checkAuths(authEnabled bool, svc util.Service, perms Permissions, auths auth.Records) Errors {
	var ret Errors

	if authEnabled {
		ret = append(ret, providerCheck(svc, &auth.ProviderGitHub, perms, auths)...)
		ret = append(ret, providerCheck(svc, &auth.ProviderGoogle, perms, auths)...)
		ret = append(ret, providerCheck(svc, &auth.ProviderSlack, perms, auths)...)
		ret = append(ret, providerCheck(svc, &auth.ProviderAmazon, perms, auths)...)
		ret = append(ret, providerCheck(svc, &auth.ProviderMicrosoft, perms, auths)...)
	}

	return ret
}

func providerCheck(svc util.Service, p *auth.Provider, perms Permissions, auths auth.Records) Errors {
	var authMatches = auths.FindByProvider(p.Key)
	var permMatches = perms.FindByK(p.Key)
	if len(permMatches) == 0 {
		return nil
	}

	var ok = false
	for _, a := range authMatches {
		for _, p := range permMatches {
			if strings.HasSuffix(a.Email, p.V) {
				ok = true
			}
		}
	}

	if ok {
		return nil
	}

	msg := "you must log in using " + p.Title

	var emailDomains []string
	for _, p := range permMatches {
		if len(p.V) > 0 {
			emailDomains = append(emailDomains, p.V)
		}
	}

	if len(emailDomains) > 0 {
		msg += " with email address " + strings.Join(emailDomains, " or ")
	}

	msg += " to access this " + svc.Key
	return Errors{{Svc: svc.Key, Provider: p.Key, Match: strings.Join(emailDomains, ","), Message: msg}}
}
