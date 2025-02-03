package web_test

import (
	"golang-integration-test/internal/app/moduleA/delivery/web"
	"golang-integration-test/internal/app/moduleA/repository"
	"golang-integration-test/internal/app/moduleA/repository/mongorepository"
	"golang-integration-test/internal/app/moduleA/usecase"
	"golang-integration-test/internal/app/moduleA/usecase/usecaseimpl"
	moduleB "golang-integration-test/internal/app/moduleB/usecase"
	moduleBImpl "golang-integration-test/internal/app/moduleB/usecase/usecaseimpl"
	"golang-integration-test/internal/pkg/persistent/mongo"
	"golang-integration-test/internal/pkg/testhelper"
	"net/http"
	"time"

	"github.com/go-chi/chi"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/tryvium-travels/memongo"
)

var mongoserver *memongo.Server

var _ = BeforeSuite(func() {
	var err error
	mongoserver, err = memongo.Start("4.0.5")
	if err != nil {
		panic(err)
	}
})

var _ = AfterSuite(func() {
	mongoserver.Stop()
})

var _ = Describe("Module A Web Handler", func() {
	var (
		moduleARepository repository.ModuleARepository
		moduleACreator    usecase.ModuleACreator
		moduleBUsecase    moduleB.ModuleBUsecase
		registry          *web.ModuleAWebRegistry
		router            chi.Router
	)

	BeforeEach(func() {
		moduleARepository = mongorepository.NewModuleAMongoRepository(mongorepository.ModuleAMongoRepositoryOptions{
			DB: mongo.ProvideMongoDb(mongoserver.URIWithRandomDB()),
		})
		moduleBUsecase = moduleBImpl.NewModuleB(moduleBImpl.ModuleBOptions{})
		moduleACreator = usecaseimpl.NewModuleACreator(usecaseimpl.ModuleACreatorOptions{
			Repository: moduleARepository,
			ModuleBUC:  moduleBUsecase,
		})
		registry = web.NewModuleAWebRegistry(web.ModuleAWebRegistryOptions{
			Creator: moduleACreator,
		})
		router = chi.NewRouter()
		registry.RegisterRoutesTo(router)
	})

	Context("POST /a/create/v1", func() {
		var (
			method  string
			uri     string
			cmd     usecase.ModuleACreatorCommand
			request *http.Request
		)

		BeforeEach(func() {
			method = http.MethodPost
			uri = "/a/create/v1"
		})

		When("request have invalid body", func() {
			BeforeEach(func() {
				request = testhelper.BuildRequest(method, uri, "invalid body")
			})

			It("should return 400 Bad Request", func() {
				result := testhelper.SendRawHTTPTestRequest(router, request)
				Expect(result.StatusCode).To(Equal(http.StatusBadRequest))
			})
		})

		When("request have valid body but with zero time given", func() {
			BeforeEach(func() {
				request = testhelper.BuildRequest(method, uri, cmd)
			})

			It("should return 400 Bad Request", func() {
				result := testhelper.SendRawHTTPTestRequest(router, request)
				Expect(result.StatusCode).To(Equal(http.StatusBadRequest))
			})
		})

		When("request have valid body but with non zero time", func() {
			BeforeEach(func() {
				cmd = usecase.ModuleACreatorCommand{
					A: time.Now(),
				}
				request = testhelper.BuildRequest(method, uri, cmd)
			})

			It("should return 200 OK", func() {
				result := testhelper.SendRawHTTPTestRequest(router, request)
				Expect(result.StatusCode).To(Equal(http.StatusOK))
			})
		})
	})
})
