package restapi

import (
	"crypto/tls"
	"net/http"

	errors "github.com/go-openapi/errors"
	runtime "github.com/go-openapi/runtime"
	middleware "github.com/go-openapi/runtime/middleware"

	"github.com/digitalrebar/digitalrebar/go/rebar-core-api/restapi/operations"
	"github.com/digitalrebar/digitalrebar/go/rebar-core-api/restapi/operations/attribs"
	"github.com/digitalrebar/digitalrebar/go/rebar-core-api/restapi/operations/availablehammers"
	"github.com/digitalrebar/digitalrebar/go/rebar-core-api/restapi/operations/barclamps"
	"github.com/digitalrebar/digitalrebar/go/rebar-core-api/restapi/operations/capabilities"
	"github.com/digitalrebar/digitalrebar/go/rebar-core-api/restapi/operations/deploymentroles"
	"github.com/digitalrebar/digitalrebar/go/rebar-core-api/restapi/operations/deployments"
	"github.com/digitalrebar/digitalrebar/go/rebar-core-api/restapi/operations/dnsnameentries"
	"github.com/digitalrebar/digitalrebar/go/rebar-core-api/restapi/operations/dnsnamefilters"
	"github.com/digitalrebar/digitalrebar/go/rebar-core-api/restapi/operations/eventselectors"
	"github.com/digitalrebar/digitalrebar/go/rebar-core-api/restapi/operations/eventsinks"
	"github.com/digitalrebar/digitalrebar/go/rebar-core-api/restapi/operations/hammers"
	"github.com/digitalrebar/digitalrebar/go/rebar-core-api/restapi/operations/jigs"
	"github.com/digitalrebar/digitalrebar/go/rebar-core-api/restapi/operations/networkallocations"
	"github.com/digitalrebar/digitalrebar/go/rebar-core-api/restapi/operations/networkranges"
	"github.com/digitalrebar/digitalrebar/go/rebar-core-api/restapi/operations/networkrouters"
	"github.com/digitalrebar/digitalrebar/go/rebar-core-api/restapi/operations/networks"
	"github.com/digitalrebar/digitalrebar/go/rebar-core-api/restapi/operations/noderoles"
	"github.com/digitalrebar/digitalrebar/go/rebar-core-api/restapi/operations/nodes"
	"github.com/digitalrebar/digitalrebar/go/rebar-core-api/restapi/operations/profiles"
	"github.com/digitalrebar/digitalrebar/go/rebar-core-api/restapi/operations/providers"
	"github.com/digitalrebar/digitalrebar/go/rebar-core-api/restapi/operations/roles"
	"github.com/digitalrebar/digitalrebar/go/rebar-core-api/restapi/operations/tenants"
	"github.com/digitalrebar/digitalrebar/go/rebar-core-api/restapi/operations/users"
	"github.com/digitalrebar/digitalrebar/go/rebar-core-api/restapi/operations/usertenantcapabilities"
)

// This file is safe to edit. Once it exists it will not be overwritten

//go:generate swagger generate server --target .. --name rebar-core-api --spec ../rebar-core.json

func configureFlags(api *operations.RebarCoreAPI) {
	// api.CommandLineOptionsGroups = []swag.CommandLineOptionsGroup{ ... }
}

func configureAPI(api *operations.RebarCoreAPI) http.Handler {
	// configure the api here
	api.ServeError = errors.ServeError

	// Set your custom logger if needed. Default one is log.Printf
	// Expected interface func(string, ...interface{})
	//
	// Example:
	// s.api.Logger = log.Printf

	api.JSONConsumer = runtime.JSONConsumer()

	api.JSONProducer = runtime.JSONProducer()

	api.AuthorizationAuth = func(token string) (interface{}, error) {
		return nil, errors.NotImplemented("api key auth (Authorization) Authorization from header has not yet been implemented")
	}

	api.AuthTokenAuth = func(token string) (interface{}, error) {
		return nil, errors.NotImplemented("api key auth (auth_token) auth_token from query has not yet been implemented")
	}

	api.AttribsDELETEAttribHandler = attribs.DELETEAttribHandlerFunc(func(params attribs.DELETEAttribParams) middleware.Responder {
		return middleware.NotImplemented("operation attribs.DELETEAttrib has not yet been implemented")
	})
	api.AvailablehammersDELETEAvailablehammersHandler = availablehammers.DELETEAvailablehammersHandlerFunc(func(params availablehammers.DELETEAvailablehammersParams) middleware.Responder {
		return middleware.NotImplemented("operation availablehammers.DELETEAvailablehammers has not yet been implemented")
	})
	api.BarclampsDELETEBarclampHandler = barclamps.DELETEBarclampHandlerFunc(func(params barclamps.DELETEBarclampParams) middleware.Responder {
		return middleware.NotImplemented("operation barclamps.DELETEBarclamp has not yet been implemented")
	})
	api.CapabilitiesDELETECapabilitiesHandler = capabilities.DELETECapabilitiesHandlerFunc(func(params capabilities.DELETECapabilitiesParams) middleware.Responder {
		return middleware.NotImplemented("operation capabilities.DELETECapabilities has not yet been implemented")
	})
	api.DeploymentrolesDELETEDeploymentrolesHandler = deploymentroles.DELETEDeploymentrolesHandlerFunc(func(params deploymentroles.DELETEDeploymentrolesParams) middleware.Responder {
		return middleware.NotImplemented("operation deploymentroles.DELETEDeploymentroles has not yet been implemented")
	})
	api.DeploymentsDELETEDeploymentsHandler = deployments.DELETEDeploymentsHandlerFunc(func(params deployments.DELETEDeploymentsParams) middleware.Responder {
		return middleware.NotImplemented("operation deployments.DELETEDeployments has not yet been implemented")
	})
	api.DnsnameentriesDELETEDnsnameentriesHandler = dnsnameentries.DELETEDnsnameentriesHandlerFunc(func(params dnsnameentries.DELETEDnsnameentriesParams) middleware.Responder {
		return middleware.NotImplemented("operation dnsnameentries.DELETEDnsnameentries has not yet been implemented")
	})
	api.DnsnamefiltersDELETEDnsnamefiltersHandler = dnsnamefilters.DELETEDnsnamefiltersHandlerFunc(func(params dnsnamefilters.DELETEDnsnamefiltersParams) middleware.Responder {
		return middleware.NotImplemented("operation dnsnamefilters.DELETEDnsnamefilters has not yet been implemented")
	})
	api.EventselectorsDELETEEventselectorHandler = eventselectors.DELETEEventselectorHandlerFunc(func(params eventselectors.DELETEEventselectorParams) middleware.Responder {
		return middleware.NotImplemented("operation eventselectors.DELETEEventselector has not yet been implemented")
	})
	api.EventsinksDELETEEventsinkHandler = eventsinks.DELETEEventsinkHandlerFunc(func(params eventsinks.DELETEEventsinkParams) middleware.Responder {
		return middleware.NotImplemented("operation eventsinks.DELETEEventsink has not yet been implemented")
	})
	api.HammersDELETEHammerHandler = hammers.DELETEHammerHandlerFunc(func(params hammers.DELETEHammerParams) middleware.Responder {
		return middleware.NotImplemented("operation hammers.DELETEHammer has not yet been implemented")
	})
	api.JigsDELETEJigHandler = jigs.DELETEJigHandlerFunc(func(params jigs.DELETEJigParams) middleware.Responder {
		return middleware.NotImplemented("operation jigs.DELETEJig has not yet been implemented")
	})
	api.NetworksDELETENetworkHandler = networks.DELETENetworkHandlerFunc(func(params networks.DELETENetworkParams) middleware.Responder {
		return middleware.NotImplemented("operation networks.DELETENetwork has not yet been implemented")
	})
	api.NetworkallocationsDELETENetworkallocationHandler = networkallocations.DELETENetworkallocationHandlerFunc(func(params networkallocations.DELETENetworkallocationParams) middleware.Responder {
		return middleware.NotImplemented("operation networkallocations.DELETENetworkallocation has not yet been implemented")
	})
	api.NetworkrangesDELETENetworkrangeHandler = networkranges.DELETENetworkrangeHandlerFunc(func(params networkranges.DELETENetworkrangeParams) middleware.Responder {
		return middleware.NotImplemented("operation networkranges.DELETENetworkrange has not yet been implemented")
	})
	api.NetworkroutersDELETENetworkrouterHandler = networkrouters.DELETENetworkrouterHandlerFunc(func(params networkrouters.DELETENetworkrouterParams) middleware.Responder {
		return middleware.NotImplemented("operation networkrouters.DELETENetworkrouter has not yet been implemented")
	})
	api.NodesDELETENodeHandler = nodes.DELETENodeHandlerFunc(func(params nodes.DELETENodeParams) middleware.Responder {
		return middleware.NotImplemented("operation nodes.DELETENode has not yet been implemented")
	})
	api.NoderolesDELETENoderoleHandler = noderoles.DELETENoderoleHandlerFunc(func(params noderoles.DELETENoderoleParams) middleware.Responder {
		return middleware.NotImplemented("operation noderoles.DELETENoderole has not yet been implemented")
	})
	api.ProfilesDELETEProfileHandler = profiles.DELETEProfileHandlerFunc(func(params profiles.DELETEProfileParams) middleware.Responder {
		return middleware.NotImplemented("operation profiles.DELETEProfile has not yet been implemented")
	})
	api.ProvidersDELETEProviderHandler = providers.DELETEProviderHandlerFunc(func(params providers.DELETEProviderParams) middleware.Responder {
		return middleware.NotImplemented("operation providers.DELETEProvider has not yet been implemented")
	})
	api.RolesDELETERolesHandler = roles.DELETERolesHandlerFunc(func(params roles.DELETERolesParams) middleware.Responder {
		return middleware.NotImplemented("operation roles.DELETERoles has not yet been implemented")
	})
	api.TenantsDELETETenantHandler = tenants.DELETETenantHandlerFunc(func(params tenants.DELETETenantParams) middleware.Responder {
		return middleware.NotImplemented("operation tenants.DELETETenant has not yet been implemented")
	})
	api.UsersDELETEUserHandler = users.DELETEUserHandlerFunc(func(params users.DELETEUserParams) middleware.Responder {
		return middleware.NotImplemented("operation users.DELETEUser has not yet been implemented")
	})
	api.UsertenantcapabilitiesDELETEUsertenantcapabilityHandler = usertenantcapabilities.DELETEUsertenantcapabilityHandlerFunc(func(params usertenantcapabilities.DELETEUsertenantcapabilityParams) middleware.Responder {
		return middleware.NotImplemented("operation usertenantcapabilities.DELETEUsertenantcapability has not yet been implemented")
	})
	api.AttribsGETAttribHandler = attribs.GETAttribHandlerFunc(func(params attribs.GETAttribParams) middleware.Responder {
		return middleware.NotImplemented("operation attribs.GETAttrib has not yet been implemented")
	})
	api.AvailablehammersGETAvailablehammersHandler = availablehammers.GETAvailablehammersHandlerFunc(func(params availablehammers.GETAvailablehammersParams) middleware.Responder {
		return middleware.NotImplemented("operation availablehammers.GETAvailablehammers has not yet been implemented")
	})
	api.BarclampsGETBarclampHandler = barclamps.GETBarclampHandlerFunc(func(params barclamps.GETBarclampParams) middleware.Responder {
		return middleware.NotImplemented("operation barclamps.GETBarclamp has not yet been implemented")
	})
	api.CapabilitiesGETCapabilitiesHandler = capabilities.GETCapabilitiesHandlerFunc(func(params capabilities.GETCapabilitiesParams) middleware.Responder {
		return middleware.NotImplemented("operation capabilities.GETCapabilities has not yet been implemented")
	})
	api.DeploymentrolesGETDeploymentrolesHandler = deploymentroles.GETDeploymentrolesHandlerFunc(func(params deploymentroles.GETDeploymentrolesParams) middleware.Responder {
		return middleware.NotImplemented("operation deploymentroles.GETDeploymentroles has not yet been implemented")
	})
	api.DeploymentsGETDeploymentsHandler = deployments.GETDeploymentsHandlerFunc(func(params deployments.GETDeploymentsParams) middleware.Responder {
		return middleware.NotImplemented("operation deployments.GETDeployments has not yet been implemented")
	})
	api.DnsnameentriesGETDnsnameentriesHandler = dnsnameentries.GETDnsnameentriesHandlerFunc(func(params dnsnameentries.GETDnsnameentriesParams) middleware.Responder {
		return middleware.NotImplemented("operation dnsnameentries.GETDnsnameentries has not yet been implemented")
	})
	api.DnsnamefiltersGETDnsnamefiltersHandler = dnsnamefilters.GETDnsnamefiltersHandlerFunc(func(params dnsnamefilters.GETDnsnamefiltersParams) middleware.Responder {
		return middleware.NotImplemented("operation dnsnamefilters.GETDnsnamefilters has not yet been implemented")
	})
	api.EventselectorsGETEventselectorHandler = eventselectors.GETEventselectorHandlerFunc(func(params eventselectors.GETEventselectorParams) middleware.Responder {
		return middleware.NotImplemented("operation eventselectors.GETEventselector has not yet been implemented")
	})
	api.EventsinksGETEventsinkHandler = eventsinks.GETEventsinkHandlerFunc(func(params eventsinks.GETEventsinkParams) middleware.Responder {
		return middleware.NotImplemented("operation eventsinks.GETEventsink has not yet been implemented")
	})
	api.HammersGETHammerHandler = hammers.GETHammerHandlerFunc(func(params hammers.GETHammerParams) middleware.Responder {
		return middleware.NotImplemented("operation hammers.GETHammer has not yet been implemented")
	})
	api.JigsGETJigHandler = jigs.GETJigHandlerFunc(func(params jigs.GETJigParams) middleware.Responder {
		return middleware.NotImplemented("operation jigs.GETJig has not yet been implemented")
	})
	api.NetworksGETNetworkHandler = networks.GETNetworkHandlerFunc(func(params networks.GETNetworkParams) middleware.Responder {
		return middleware.NotImplemented("operation networks.GETNetwork has not yet been implemented")
	})
	api.NetworkallocationsGETNetworkallocationHandler = networkallocations.GETNetworkallocationHandlerFunc(func(params networkallocations.GETNetworkallocationParams) middleware.Responder {
		return middleware.NotImplemented("operation networkallocations.GETNetworkallocation has not yet been implemented")
	})
	api.NetworkrangesGETNetworkrangeHandler = networkranges.GETNetworkrangeHandlerFunc(func(params networkranges.GETNetworkrangeParams) middleware.Responder {
		return middleware.NotImplemented("operation networkranges.GETNetworkrange has not yet been implemented")
	})
	api.NetworkroutersGETNetworkrouterHandler = networkrouters.GETNetworkrouterHandlerFunc(func(params networkrouters.GETNetworkrouterParams) middleware.Responder {
		return middleware.NotImplemented("operation networkrouters.GETNetworkrouter has not yet been implemented")
	})
	api.NodesGETNodeHandler = nodes.GETNodeHandlerFunc(func(params nodes.GETNodeParams, principal interface{}) middleware.Responder {
		return middleware.NotImplemented("operation nodes.GETNode has not yet been implemented")
	})
	api.NoderolesGETNoderoleHandler = noderoles.GETNoderoleHandlerFunc(func(params noderoles.GETNoderoleParams) middleware.Responder {
		return middleware.NotImplemented("operation noderoles.GETNoderole has not yet been implemented")
	})
	api.ProfilesGETProfileHandler = profiles.GETProfileHandlerFunc(func(params profiles.GETProfileParams) middleware.Responder {
		return middleware.NotImplemented("operation profiles.GETProfile has not yet been implemented")
	})
	api.ProvidersGETProviderHandler = providers.GETProviderHandlerFunc(func(params providers.GETProviderParams) middleware.Responder {
		return middleware.NotImplemented("operation providers.GETProvider has not yet been implemented")
	})
	api.RolesGETRolesHandler = roles.GETRolesHandlerFunc(func(params roles.GETRolesParams) middleware.Responder {
		return middleware.NotImplemented("operation roles.GETRoles has not yet been implemented")
	})
	api.TenantsGETTenantHandler = tenants.GETTenantHandlerFunc(func(params tenants.GETTenantParams) middleware.Responder {
		return middleware.NotImplemented("operation tenants.GETTenant has not yet been implemented")
	})
	api.UsersGETUserHandler = users.GETUserHandlerFunc(func(params users.GETUserParams) middleware.Responder {
		return middleware.NotImplemented("operation users.GETUser has not yet been implemented")
	})
	api.UsertenantcapabilitiesGETUsertenantcapabilityHandler = usertenantcapabilities.GETUsertenantcapabilityHandlerFunc(func(params usertenantcapabilities.GETUsertenantcapabilityParams) middleware.Responder {
		return middleware.NotImplemented("operation usertenantcapabilities.GETUsertenantcapability has not yet been implemented")
	})
	api.AttribsLISTAttribsHandler = attribs.LISTAttribsHandlerFunc(func(params attribs.LISTAttribsParams) middleware.Responder {
		return middleware.NotImplemented("operation attribs.LISTAttribs has not yet been implemented")
	})
	api.AvailablehammersLISTAvailablehammersHandler = availablehammers.LISTAvailablehammersHandlerFunc(func(params availablehammers.LISTAvailablehammersParams) middleware.Responder {
		return middleware.NotImplemented("operation availablehammers.LISTAvailablehammers has not yet been implemented")
	})
	api.BarclampsLISTBarclampsHandler = barclamps.LISTBarclampsHandlerFunc(func(params barclamps.LISTBarclampsParams) middleware.Responder {
		return middleware.NotImplemented("operation barclamps.LISTBarclamps has not yet been implemented")
	})
	api.CapabilitiesLISTCapabilitiesHandler = capabilities.LISTCapabilitiesHandlerFunc(func(params capabilities.LISTCapabilitiesParams) middleware.Responder {
		return middleware.NotImplemented("operation capabilities.LISTCapabilities has not yet been implemented")
	})
	api.DeploymentrolesLISTDeploymentrolesHandler = deploymentroles.LISTDeploymentrolesHandlerFunc(func(params deploymentroles.LISTDeploymentrolesParams) middleware.Responder {
		return middleware.NotImplemented("operation deploymentroles.LISTDeploymentroles has not yet been implemented")
	})
	api.DeploymentsLISTDeploymentsHandler = deployments.LISTDeploymentsHandlerFunc(func(params deployments.LISTDeploymentsParams) middleware.Responder {
		return middleware.NotImplemented("operation deployments.LISTDeployments has not yet been implemented")
	})
	api.DnsnameentriesLISTDnsnameentriesHandler = dnsnameentries.LISTDnsnameentriesHandlerFunc(func(params dnsnameentries.LISTDnsnameentriesParams) middleware.Responder {
		return middleware.NotImplemented("operation dnsnameentries.LISTDnsnameentries has not yet been implemented")
	})
	api.DnsnamefiltersLISTDnsnamefiltersHandler = dnsnamefilters.LISTDnsnamefiltersHandlerFunc(func(params dnsnamefilters.LISTDnsnamefiltersParams) middleware.Responder {
		return middleware.NotImplemented("operation dnsnamefilters.LISTDnsnamefilters has not yet been implemented")
	})
	api.EventselectorsLISTEventselectorsHandler = eventselectors.LISTEventselectorsHandlerFunc(func(params eventselectors.LISTEventselectorsParams) middleware.Responder {
		return middleware.NotImplemented("operation eventselectors.LISTEventselectors has not yet been implemented")
	})
	api.EventsinksLISTEventsinksHandler = eventsinks.LISTEventsinksHandlerFunc(func(params eventsinks.LISTEventsinksParams) middleware.Responder {
		return middleware.NotImplemented("operation eventsinks.LISTEventsinks has not yet been implemented")
	})
	api.HammersLISTHammersHandler = hammers.LISTHammersHandlerFunc(func(params hammers.LISTHammersParams) middleware.Responder {
		return middleware.NotImplemented("operation hammers.LISTHammers has not yet been implemented")
	})
	api.JigsLISTJigsHandler = jigs.LISTJigsHandlerFunc(func(params jigs.LISTJigsParams) middleware.Responder {
		return middleware.NotImplemented("operation jigs.LISTJigs has not yet been implemented")
	})
	api.NetworkallocationsLISTNetworkallocationsHandler = networkallocations.LISTNetworkallocationsHandlerFunc(func(params networkallocations.LISTNetworkallocationsParams) middleware.Responder {
		return middleware.NotImplemented("operation networkallocations.LISTNetworkallocations has not yet been implemented")
	})
	api.NetworkrangesLISTNetworkrangesHandler = networkranges.LISTNetworkrangesHandlerFunc(func(params networkranges.LISTNetworkrangesParams) middleware.Responder {
		return middleware.NotImplemented("operation networkranges.LISTNetworkranges has not yet been implemented")
	})
	api.NetworkroutersLISTNetworkroutersHandler = networkrouters.LISTNetworkroutersHandlerFunc(func(params networkrouters.LISTNetworkroutersParams) middleware.Responder {
		return middleware.NotImplemented("operation networkrouters.LISTNetworkrouters has not yet been implemented")
	})
	api.NetworksLISTNetworksHandler = networks.LISTNetworksHandlerFunc(func(params networks.LISTNetworksParams) middleware.Responder {
		return middleware.NotImplemented("operation networks.LISTNetworks has not yet been implemented")
	})
	api.NoderolesLISTNoderolesHandler = noderoles.LISTNoderolesHandlerFunc(func(params noderoles.LISTNoderolesParams) middleware.Responder {
		return middleware.NotImplemented("operation noderoles.LISTNoderoles has not yet been implemented")
	})
	api.NodesLISTNodesHandler = nodes.LISTNodesHandlerFunc(func(params nodes.LISTNodesParams) middleware.Responder {
		return middleware.NotImplemented("operation nodes.LISTNodes has not yet been implemented")
	})
	api.ProfilesLISTProfilesHandler = profiles.LISTProfilesHandlerFunc(func(params profiles.LISTProfilesParams) middleware.Responder {
		return middleware.NotImplemented("operation profiles.LISTProfiles has not yet been implemented")
	})
	api.ProvidersLISTProvidersHandler = providers.LISTProvidersHandlerFunc(func(params providers.LISTProvidersParams) middleware.Responder {
		return middleware.NotImplemented("operation providers.LISTProviders has not yet been implemented")
	})
	api.RolesLISTRolesHandler = roles.LISTRolesHandlerFunc(func(params roles.LISTRolesParams) middleware.Responder {
		return middleware.NotImplemented("operation roles.LISTRoles has not yet been implemented")
	})
	api.TenantsLISTTenantsHandler = tenants.LISTTenantsHandlerFunc(func(params tenants.LISTTenantsParams) middleware.Responder {
		return middleware.NotImplemented("operation tenants.LISTTenants has not yet been implemented")
	})
	api.UsersLISTUsersHandler = users.LISTUsersHandlerFunc(func(params users.LISTUsersParams) middleware.Responder {
		return middleware.NotImplemented("operation users.LISTUsers has not yet been implemented")
	})
	api.UsertenantcapabilitiesLISTUsertenantcapabilitiesHandler = usertenantcapabilities.LISTUsertenantcapabilitiesHandlerFunc(func(params usertenantcapabilities.LISTUsertenantcapabilitiesParams) middleware.Responder {
		return middleware.NotImplemented("operation usertenantcapabilities.LISTUsertenantcapabilities has not yet been implemented")
	})
	api.NodesPATCHNodeHandler = nodes.PATCHNodeHandlerFunc(func(params nodes.PATCHNodeParams) middleware.Responder {
		return middleware.NotImplemented("operation nodes.PATCHNode has not yet been implemented")
	})
	api.AttribsPOSTAttribHandler = attribs.POSTAttribHandlerFunc(func(params attribs.POSTAttribParams) middleware.Responder {
		return middleware.NotImplemented("operation attribs.POSTAttrib has not yet been implemented")
	})
	api.AvailablehammersPOSTAvailablehammersHandler = availablehammers.POSTAvailablehammersHandlerFunc(func(params availablehammers.POSTAvailablehammersParams) middleware.Responder {
		return middleware.NotImplemented("operation availablehammers.POSTAvailablehammers has not yet been implemented")
	})
	api.BarclampsPOSTBarclampHandler = barclamps.POSTBarclampHandlerFunc(func(params barclamps.POSTBarclampParams) middleware.Responder {
		return middleware.NotImplemented("operation barclamps.POSTBarclamp has not yet been implemented")
	})
	api.CapabilitiesPOSTCapabilitiesHandler = capabilities.POSTCapabilitiesHandlerFunc(func(params capabilities.POSTCapabilitiesParams) middleware.Responder {
		return middleware.NotImplemented("operation capabilities.POSTCapabilities has not yet been implemented")
	})
	api.DeploymentrolesPOSTDeploymentrolesHandler = deploymentroles.POSTDeploymentrolesHandlerFunc(func(params deploymentroles.POSTDeploymentrolesParams) middleware.Responder {
		return middleware.NotImplemented("operation deploymentroles.POSTDeploymentroles has not yet been implemented")
	})
	api.DeploymentsPOSTDeploymentsHandler = deployments.POSTDeploymentsHandlerFunc(func(params deployments.POSTDeploymentsParams) middleware.Responder {
		return middleware.NotImplemented("operation deployments.POSTDeployments has not yet been implemented")
	})
	api.DnsnameentriesPOSTDnsnameentriesHandler = dnsnameentries.POSTDnsnameentriesHandlerFunc(func(params dnsnameentries.POSTDnsnameentriesParams) middleware.Responder {
		return middleware.NotImplemented("operation dnsnameentries.POSTDnsnameentries has not yet been implemented")
	})
	api.DnsnamefiltersPOSTDnsnamefiltersHandler = dnsnamefilters.POSTDnsnamefiltersHandlerFunc(func(params dnsnamefilters.POSTDnsnamefiltersParams) middleware.Responder {
		return middleware.NotImplemented("operation dnsnamefilters.POSTDnsnamefilters has not yet been implemented")
	})
	api.EventselectorsPOSTEventselectorHandler = eventselectors.POSTEventselectorHandlerFunc(func(params eventselectors.POSTEventselectorParams) middleware.Responder {
		return middleware.NotImplemented("operation eventselectors.POSTEventselector has not yet been implemented")
	})
	api.EventsinksPOSTEventsinkHandler = eventsinks.POSTEventsinkHandlerFunc(func(params eventsinks.POSTEventsinkParams) middleware.Responder {
		return middleware.NotImplemented("operation eventsinks.POSTEventsink has not yet been implemented")
	})
	api.HammersPOSTHammerHandler = hammers.POSTHammerHandlerFunc(func(params hammers.POSTHammerParams) middleware.Responder {
		return middleware.NotImplemented("operation hammers.POSTHammer has not yet been implemented")
	})
	api.JigsPOSTJigHandler = jigs.POSTJigHandlerFunc(func(params jigs.POSTJigParams) middleware.Responder {
		return middleware.NotImplemented("operation jigs.POSTJig has not yet been implemented")
	})
	api.NetworksPOSTNetworkHandler = networks.POSTNetworkHandlerFunc(func(params networks.POSTNetworkParams) middleware.Responder {
		return middleware.NotImplemented("operation networks.POSTNetwork has not yet been implemented")
	})
	api.NetworkallocationsPOSTNetworkallocationHandler = networkallocations.POSTNetworkallocationHandlerFunc(func(params networkallocations.POSTNetworkallocationParams) middleware.Responder {
		return middleware.NotImplemented("operation networkallocations.POSTNetworkallocation has not yet been implemented")
	})
	api.NetworkrangesPOSTNetworkrangeHandler = networkranges.POSTNetworkrangeHandlerFunc(func(params networkranges.POSTNetworkrangeParams) middleware.Responder {
		return middleware.NotImplemented("operation networkranges.POSTNetworkrange has not yet been implemented")
	})
	api.NetworkroutersPOSTNetworkrouterHandler = networkrouters.POSTNetworkrouterHandlerFunc(func(params networkrouters.POSTNetworkrouterParams) middleware.Responder {
		return middleware.NotImplemented("operation networkrouters.POSTNetworkrouter has not yet been implemented")
	})
	api.NodesPOSTNodeHandler = nodes.POSTNodeHandlerFunc(func(params nodes.POSTNodeParams) middleware.Responder {
		return middleware.NotImplemented("operation nodes.POSTNode has not yet been implemented")
	})
	api.NoderolesPOSTNoderoleHandler = noderoles.POSTNoderoleHandlerFunc(func(params noderoles.POSTNoderoleParams) middleware.Responder {
		return middleware.NotImplemented("operation noderoles.POSTNoderole has not yet been implemented")
	})
	api.ProfilesPOSTProfileHandler = profiles.POSTProfileHandlerFunc(func(params profiles.POSTProfileParams) middleware.Responder {
		return middleware.NotImplemented("operation profiles.POSTProfile has not yet been implemented")
	})
	api.ProvidersPOSTProviderHandler = providers.POSTProviderHandlerFunc(func(params providers.POSTProviderParams) middleware.Responder {
		return middleware.NotImplemented("operation providers.POSTProvider has not yet been implemented")
	})
	api.RolesPOSTRolesHandler = roles.POSTRolesHandlerFunc(func(params roles.POSTRolesParams) middleware.Responder {
		return middleware.NotImplemented("operation roles.POSTRoles has not yet been implemented")
	})
	api.TenantsPOSTTenantHandler = tenants.POSTTenantHandlerFunc(func(params tenants.POSTTenantParams) middleware.Responder {
		return middleware.NotImplemented("operation tenants.POSTTenant has not yet been implemented")
	})
	api.UsersPOSTUserHandler = users.POSTUserHandlerFunc(func(params users.POSTUserParams) middleware.Responder {
		return middleware.NotImplemented("operation users.POSTUser has not yet been implemented")
	})
	api.UsertenantcapabilitiesPOSTUsertenantcapabilityHandler = usertenantcapabilities.POSTUsertenantcapabilityHandlerFunc(func(params usertenantcapabilities.POSTUsertenantcapabilityParams) middleware.Responder {
		return middleware.NotImplemented("operation usertenantcapabilities.POSTUsertenantcapability has not yet been implemented")
	})
	api.AttribsPUTAttribHandler = attribs.PUTAttribHandlerFunc(func(params attribs.PUTAttribParams) middleware.Responder {
		return middleware.NotImplemented("operation attribs.PUTAttrib has not yet been implemented")
	})
	api.AvailablehammersPUTAvailablehammersHandler = availablehammers.PUTAvailablehammersHandlerFunc(func(params availablehammers.PUTAvailablehammersParams) middleware.Responder {
		return middleware.NotImplemented("operation availablehammers.PUTAvailablehammers has not yet been implemented")
	})
	api.BarclampsPUTBarclampHandler = barclamps.PUTBarclampHandlerFunc(func(params barclamps.PUTBarclampParams) middleware.Responder {
		return middleware.NotImplemented("operation barclamps.PUTBarclamp has not yet been implemented")
	})
	api.CapabilitiesPUTCapabilitiesHandler = capabilities.PUTCapabilitiesHandlerFunc(func(params capabilities.PUTCapabilitiesParams) middleware.Responder {
		return middleware.NotImplemented("operation capabilities.PUTCapabilities has not yet been implemented")
	})
	api.DeploymentrolesPUTDeploymentrolesHandler = deploymentroles.PUTDeploymentrolesHandlerFunc(func(params deploymentroles.PUTDeploymentrolesParams) middleware.Responder {
		return middleware.NotImplemented("operation deploymentroles.PUTDeploymentroles has not yet been implemented")
	})
	api.DeploymentsPUTDeploymentsHandler = deployments.PUTDeploymentsHandlerFunc(func(params deployments.PUTDeploymentsParams) middleware.Responder {
		return middleware.NotImplemented("operation deployments.PUTDeployments has not yet been implemented")
	})
	api.DnsnameentriesPUTDnsnameentriesHandler = dnsnameentries.PUTDnsnameentriesHandlerFunc(func(params dnsnameentries.PUTDnsnameentriesParams) middleware.Responder {
		return middleware.NotImplemented("operation dnsnameentries.PUTDnsnameentries has not yet been implemented")
	})
	api.DnsnamefiltersPUTDnsnamefiltersHandler = dnsnamefilters.PUTDnsnamefiltersHandlerFunc(func(params dnsnamefilters.PUTDnsnamefiltersParams) middleware.Responder {
		return middleware.NotImplemented("operation dnsnamefilters.PUTDnsnamefilters has not yet been implemented")
	})
	api.EventselectorsPUTEventselectorHandler = eventselectors.PUTEventselectorHandlerFunc(func(params eventselectors.PUTEventselectorParams) middleware.Responder {
		return middleware.NotImplemented("operation eventselectors.PUTEventselector has not yet been implemented")
	})
	api.EventsinksPUTEventsinkHandler = eventsinks.PUTEventsinkHandlerFunc(func(params eventsinks.PUTEventsinkParams) middleware.Responder {
		return middleware.NotImplemented("operation eventsinks.PUTEventsink has not yet been implemented")
	})
	api.HammersPUTHammerHandler = hammers.PUTHammerHandlerFunc(func(params hammers.PUTHammerParams) middleware.Responder {
		return middleware.NotImplemented("operation hammers.PUTHammer has not yet been implemented")
	})
	api.JigsPUTJigHandler = jigs.PUTJigHandlerFunc(func(params jigs.PUTJigParams) middleware.Responder {
		return middleware.NotImplemented("operation jigs.PUTJig has not yet been implemented")
	})
	api.NetworksPUTNetworkHandler = networks.PUTNetworkHandlerFunc(func(params networks.PUTNetworkParams) middleware.Responder {
		return middleware.NotImplemented("operation networks.PUTNetwork has not yet been implemented")
	})
	api.NetworkallocationsPUTNetworkallocationHandler = networkallocations.PUTNetworkallocationHandlerFunc(func(params networkallocations.PUTNetworkallocationParams) middleware.Responder {
		return middleware.NotImplemented("operation networkallocations.PUTNetworkallocation has not yet been implemented")
	})
	api.NetworkrangesPUTNetworkrangeHandler = networkranges.PUTNetworkrangeHandlerFunc(func(params networkranges.PUTNetworkrangeParams) middleware.Responder {
		return middleware.NotImplemented("operation networkranges.PUTNetworkrange has not yet been implemented")
	})
	api.NetworkroutersPUTNetworkrouterHandler = networkrouters.PUTNetworkrouterHandlerFunc(func(params networkrouters.PUTNetworkrouterParams) middleware.Responder {
		return middleware.NotImplemented("operation networkrouters.PUTNetworkrouter has not yet been implemented")
	})
	api.NodesPUTNodeHandler = nodes.PUTNodeHandlerFunc(func(params nodes.PUTNodeParams) middleware.Responder {
		return middleware.NotImplemented("operation nodes.PUTNode has not yet been implemented")
	})
	api.NoderolesPUTNoderoleHandler = noderoles.PUTNoderoleHandlerFunc(func(params noderoles.PUTNoderoleParams) middleware.Responder {
		return middleware.NotImplemented("operation noderoles.PUTNoderole has not yet been implemented")
	})
	api.ProfilesPUTProfileHandler = profiles.PUTProfileHandlerFunc(func(params profiles.PUTProfileParams) middleware.Responder {
		return middleware.NotImplemented("operation profiles.PUTProfile has not yet been implemented")
	})
	api.ProvidersPUTProviderHandler = providers.PUTProviderHandlerFunc(func(params providers.PUTProviderParams) middleware.Responder {
		return middleware.NotImplemented("operation providers.PUTProvider has not yet been implemented")
	})
	api.RolesPUTRolesHandler = roles.PUTRolesHandlerFunc(func(params roles.PUTRolesParams) middleware.Responder {
		return middleware.NotImplemented("operation roles.PUTRoles has not yet been implemented")
	})
	api.TenantsPUTTenantHandler = tenants.PUTTenantHandlerFunc(func(params tenants.PUTTenantParams) middleware.Responder {
		return middleware.NotImplemented("operation tenants.PUTTenant has not yet been implemented")
	})
	api.UsersPUTUserHandler = users.PUTUserHandlerFunc(func(params users.PUTUserParams) middleware.Responder {
		return middleware.NotImplemented("operation users.PUTUser has not yet been implemented")
	})
	api.UsertenantcapabilitiesPUTUsertenantcapabilityHandler = usertenantcapabilities.PUTUsertenantcapabilityHandlerFunc(func(params usertenantcapabilities.PUTUsertenantcapabilityParams) middleware.Responder {
		return middleware.NotImplemented("operation usertenantcapabilities.PUTUsertenantcapability has not yet been implemented")
	})

	api.ServerShutdown = func() {}

	return setupGlobalMiddleware(api.Serve(setupMiddlewares))
}

// The TLS configuration before HTTPS server starts.
func configureTLS(tlsConfig *tls.Config) {
	// Make all necessary changes to the TLS configuration here.
}

// The middleware configuration is for the handler executors. These do not apply to the swagger.json document.
// The middleware executes after routing but before authentication, binding and validation
func setupMiddlewares(handler http.Handler) http.Handler {
	return handler
}

// The middleware configuration happens before anything, this middleware also applies to serving the swagger.json document.
// So this is a good place to plug in a panic handling middleware, logging and metrics
func setupGlobalMiddleware(handler http.Handler) http.Handler {
	return handler
}
