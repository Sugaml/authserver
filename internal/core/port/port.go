package port

type IService interface {
	ApplicationService
	// TokenService
	ClientService
	CustomerService
	ResourceService
	RoleService
	ClientSecretService
	TenantService
	UserService
}
