/*
Copyright AppsCode Inc. and Contributors

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by client-gen. DO NOT EDIT.

package versioned

import (
	"fmt"

	addonv1alpha1 "kubeform.dev/provider-pagerduty-api/client/clientset/versioned/typed/addon/v1alpha1"
	businessv1alpha1 "kubeform.dev/provider-pagerduty-api/client/clientset/versioned/typed/business/v1alpha1"
	escalationv1alpha1 "kubeform.dev/provider-pagerduty-api/client/clientset/versioned/typed/escalation/v1alpha1"
	eventv1alpha1 "kubeform.dev/provider-pagerduty-api/client/clientset/versioned/typed/event/v1alpha1"
	extensionv1alpha1 "kubeform.dev/provider-pagerduty-api/client/clientset/versioned/typed/extension/v1alpha1"
	maintenancev1alpha1 "kubeform.dev/provider-pagerduty-api/client/clientset/versioned/typed/maintenance/v1alpha1"
	responsev1alpha1 "kubeform.dev/provider-pagerduty-api/client/clientset/versioned/typed/response/v1alpha1"
	rulesetv1alpha1 "kubeform.dev/provider-pagerduty-api/client/clientset/versioned/typed/ruleset/v1alpha1"
	schedulev1alpha1 "kubeform.dev/provider-pagerduty-api/client/clientset/versioned/typed/schedule/v1alpha1"
	servicev1alpha1 "kubeform.dev/provider-pagerduty-api/client/clientset/versioned/typed/service/v1alpha1"
	slackv1alpha1 "kubeform.dev/provider-pagerduty-api/client/clientset/versioned/typed/slack/v1alpha1"
	tagv1alpha1 "kubeform.dev/provider-pagerduty-api/client/clientset/versioned/typed/tag/v1alpha1"
	teamv1alpha1 "kubeform.dev/provider-pagerduty-api/client/clientset/versioned/typed/team/v1alpha1"
	userv1alpha1 "kubeform.dev/provider-pagerduty-api/client/clientset/versioned/typed/user/v1alpha1"

	discovery "k8s.io/client-go/discovery"
	rest "k8s.io/client-go/rest"
	flowcontrol "k8s.io/client-go/util/flowcontrol"
)

type Interface interface {
	Discovery() discovery.DiscoveryInterface
	AddonV1alpha1() addonv1alpha1.AddonV1alpha1Interface
	BusinessV1alpha1() businessv1alpha1.BusinessV1alpha1Interface
	EscalationV1alpha1() escalationv1alpha1.EscalationV1alpha1Interface
	EventV1alpha1() eventv1alpha1.EventV1alpha1Interface
	ExtensionV1alpha1() extensionv1alpha1.ExtensionV1alpha1Interface
	MaintenanceV1alpha1() maintenancev1alpha1.MaintenanceV1alpha1Interface
	ResponseV1alpha1() responsev1alpha1.ResponseV1alpha1Interface
	RulesetV1alpha1() rulesetv1alpha1.RulesetV1alpha1Interface
	ScheduleV1alpha1() schedulev1alpha1.ScheduleV1alpha1Interface
	ServiceV1alpha1() servicev1alpha1.ServiceV1alpha1Interface
	SlackV1alpha1() slackv1alpha1.SlackV1alpha1Interface
	TagV1alpha1() tagv1alpha1.TagV1alpha1Interface
	TeamV1alpha1() teamv1alpha1.TeamV1alpha1Interface
	UserV1alpha1() userv1alpha1.UserV1alpha1Interface
}

// Clientset contains the clients for groups. Each group has exactly one
// version included in a Clientset.
type Clientset struct {
	*discovery.DiscoveryClient
	addonV1alpha1       *addonv1alpha1.AddonV1alpha1Client
	businessV1alpha1    *businessv1alpha1.BusinessV1alpha1Client
	escalationV1alpha1  *escalationv1alpha1.EscalationV1alpha1Client
	eventV1alpha1       *eventv1alpha1.EventV1alpha1Client
	extensionV1alpha1   *extensionv1alpha1.ExtensionV1alpha1Client
	maintenanceV1alpha1 *maintenancev1alpha1.MaintenanceV1alpha1Client
	responseV1alpha1    *responsev1alpha1.ResponseV1alpha1Client
	rulesetV1alpha1     *rulesetv1alpha1.RulesetV1alpha1Client
	scheduleV1alpha1    *schedulev1alpha1.ScheduleV1alpha1Client
	serviceV1alpha1     *servicev1alpha1.ServiceV1alpha1Client
	slackV1alpha1       *slackv1alpha1.SlackV1alpha1Client
	tagV1alpha1         *tagv1alpha1.TagV1alpha1Client
	teamV1alpha1        *teamv1alpha1.TeamV1alpha1Client
	userV1alpha1        *userv1alpha1.UserV1alpha1Client
}

// AddonV1alpha1 retrieves the AddonV1alpha1Client
func (c *Clientset) AddonV1alpha1() addonv1alpha1.AddonV1alpha1Interface {
	return c.addonV1alpha1
}

// BusinessV1alpha1 retrieves the BusinessV1alpha1Client
func (c *Clientset) BusinessV1alpha1() businessv1alpha1.BusinessV1alpha1Interface {
	return c.businessV1alpha1
}

// EscalationV1alpha1 retrieves the EscalationV1alpha1Client
func (c *Clientset) EscalationV1alpha1() escalationv1alpha1.EscalationV1alpha1Interface {
	return c.escalationV1alpha1
}

// EventV1alpha1 retrieves the EventV1alpha1Client
func (c *Clientset) EventV1alpha1() eventv1alpha1.EventV1alpha1Interface {
	return c.eventV1alpha1
}

// ExtensionV1alpha1 retrieves the ExtensionV1alpha1Client
func (c *Clientset) ExtensionV1alpha1() extensionv1alpha1.ExtensionV1alpha1Interface {
	return c.extensionV1alpha1
}

// MaintenanceV1alpha1 retrieves the MaintenanceV1alpha1Client
func (c *Clientset) MaintenanceV1alpha1() maintenancev1alpha1.MaintenanceV1alpha1Interface {
	return c.maintenanceV1alpha1
}

// ResponseV1alpha1 retrieves the ResponseV1alpha1Client
func (c *Clientset) ResponseV1alpha1() responsev1alpha1.ResponseV1alpha1Interface {
	return c.responseV1alpha1
}

// RulesetV1alpha1 retrieves the RulesetV1alpha1Client
func (c *Clientset) RulesetV1alpha1() rulesetv1alpha1.RulesetV1alpha1Interface {
	return c.rulesetV1alpha1
}

// ScheduleV1alpha1 retrieves the ScheduleV1alpha1Client
func (c *Clientset) ScheduleV1alpha1() schedulev1alpha1.ScheduleV1alpha1Interface {
	return c.scheduleV1alpha1
}

// ServiceV1alpha1 retrieves the ServiceV1alpha1Client
func (c *Clientset) ServiceV1alpha1() servicev1alpha1.ServiceV1alpha1Interface {
	return c.serviceV1alpha1
}

// SlackV1alpha1 retrieves the SlackV1alpha1Client
func (c *Clientset) SlackV1alpha1() slackv1alpha1.SlackV1alpha1Interface {
	return c.slackV1alpha1
}

// TagV1alpha1 retrieves the TagV1alpha1Client
func (c *Clientset) TagV1alpha1() tagv1alpha1.TagV1alpha1Interface {
	return c.tagV1alpha1
}

// TeamV1alpha1 retrieves the TeamV1alpha1Client
func (c *Clientset) TeamV1alpha1() teamv1alpha1.TeamV1alpha1Interface {
	return c.teamV1alpha1
}

// UserV1alpha1 retrieves the UserV1alpha1Client
func (c *Clientset) UserV1alpha1() userv1alpha1.UserV1alpha1Interface {
	return c.userV1alpha1
}

// Discovery retrieves the DiscoveryClient
func (c *Clientset) Discovery() discovery.DiscoveryInterface {
	if c == nil {
		return nil
	}
	return c.DiscoveryClient
}

// NewForConfig creates a new Clientset for the given config.
// If config's RateLimiter is not set and QPS and Burst are acceptable,
// NewForConfig will generate a rate-limiter in configShallowCopy.
func NewForConfig(c *rest.Config) (*Clientset, error) {
	configShallowCopy := *c
	if configShallowCopy.RateLimiter == nil && configShallowCopy.QPS > 0 {
		if configShallowCopy.Burst <= 0 {
			return nil, fmt.Errorf("burst is required to be greater than 0 when RateLimiter is not set and QPS is set to greater than 0")
		}
		configShallowCopy.RateLimiter = flowcontrol.NewTokenBucketRateLimiter(configShallowCopy.QPS, configShallowCopy.Burst)
	}
	var cs Clientset
	var err error
	cs.addonV1alpha1, err = addonv1alpha1.NewForConfig(&configShallowCopy)
	if err != nil {
		return nil, err
	}
	cs.businessV1alpha1, err = businessv1alpha1.NewForConfig(&configShallowCopy)
	if err != nil {
		return nil, err
	}
	cs.escalationV1alpha1, err = escalationv1alpha1.NewForConfig(&configShallowCopy)
	if err != nil {
		return nil, err
	}
	cs.eventV1alpha1, err = eventv1alpha1.NewForConfig(&configShallowCopy)
	if err != nil {
		return nil, err
	}
	cs.extensionV1alpha1, err = extensionv1alpha1.NewForConfig(&configShallowCopy)
	if err != nil {
		return nil, err
	}
	cs.maintenanceV1alpha1, err = maintenancev1alpha1.NewForConfig(&configShallowCopy)
	if err != nil {
		return nil, err
	}
	cs.responseV1alpha1, err = responsev1alpha1.NewForConfig(&configShallowCopy)
	if err != nil {
		return nil, err
	}
	cs.rulesetV1alpha1, err = rulesetv1alpha1.NewForConfig(&configShallowCopy)
	if err != nil {
		return nil, err
	}
	cs.scheduleV1alpha1, err = schedulev1alpha1.NewForConfig(&configShallowCopy)
	if err != nil {
		return nil, err
	}
	cs.serviceV1alpha1, err = servicev1alpha1.NewForConfig(&configShallowCopy)
	if err != nil {
		return nil, err
	}
	cs.slackV1alpha1, err = slackv1alpha1.NewForConfig(&configShallowCopy)
	if err != nil {
		return nil, err
	}
	cs.tagV1alpha1, err = tagv1alpha1.NewForConfig(&configShallowCopy)
	if err != nil {
		return nil, err
	}
	cs.teamV1alpha1, err = teamv1alpha1.NewForConfig(&configShallowCopy)
	if err != nil {
		return nil, err
	}
	cs.userV1alpha1, err = userv1alpha1.NewForConfig(&configShallowCopy)
	if err != nil {
		return nil, err
	}

	cs.DiscoveryClient, err = discovery.NewDiscoveryClientForConfig(&configShallowCopy)
	if err != nil {
		return nil, err
	}
	return &cs, nil
}

// NewForConfigOrDie creates a new Clientset for the given config and
// panics if there is an error in the config.
func NewForConfigOrDie(c *rest.Config) *Clientset {
	var cs Clientset
	cs.addonV1alpha1 = addonv1alpha1.NewForConfigOrDie(c)
	cs.businessV1alpha1 = businessv1alpha1.NewForConfigOrDie(c)
	cs.escalationV1alpha1 = escalationv1alpha1.NewForConfigOrDie(c)
	cs.eventV1alpha1 = eventv1alpha1.NewForConfigOrDie(c)
	cs.extensionV1alpha1 = extensionv1alpha1.NewForConfigOrDie(c)
	cs.maintenanceV1alpha1 = maintenancev1alpha1.NewForConfigOrDie(c)
	cs.responseV1alpha1 = responsev1alpha1.NewForConfigOrDie(c)
	cs.rulesetV1alpha1 = rulesetv1alpha1.NewForConfigOrDie(c)
	cs.scheduleV1alpha1 = schedulev1alpha1.NewForConfigOrDie(c)
	cs.serviceV1alpha1 = servicev1alpha1.NewForConfigOrDie(c)
	cs.slackV1alpha1 = slackv1alpha1.NewForConfigOrDie(c)
	cs.tagV1alpha1 = tagv1alpha1.NewForConfigOrDie(c)
	cs.teamV1alpha1 = teamv1alpha1.NewForConfigOrDie(c)
	cs.userV1alpha1 = userv1alpha1.NewForConfigOrDie(c)

	cs.DiscoveryClient = discovery.NewDiscoveryClientForConfigOrDie(c)
	return &cs
}

// New creates a new Clientset for the given RESTClient.
func New(c rest.Interface) *Clientset {
	var cs Clientset
	cs.addonV1alpha1 = addonv1alpha1.New(c)
	cs.businessV1alpha1 = businessv1alpha1.New(c)
	cs.escalationV1alpha1 = escalationv1alpha1.New(c)
	cs.eventV1alpha1 = eventv1alpha1.New(c)
	cs.extensionV1alpha1 = extensionv1alpha1.New(c)
	cs.maintenanceV1alpha1 = maintenancev1alpha1.New(c)
	cs.responseV1alpha1 = responsev1alpha1.New(c)
	cs.rulesetV1alpha1 = rulesetv1alpha1.New(c)
	cs.scheduleV1alpha1 = schedulev1alpha1.New(c)
	cs.serviceV1alpha1 = servicev1alpha1.New(c)
	cs.slackV1alpha1 = slackv1alpha1.New(c)
	cs.tagV1alpha1 = tagv1alpha1.New(c)
	cs.teamV1alpha1 = teamv1alpha1.New(c)
	cs.userV1alpha1 = userv1alpha1.New(c)

	cs.DiscoveryClient = discovery.NewDiscoveryClient(c)
	return &cs
}
