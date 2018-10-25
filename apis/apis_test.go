package apis_test

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/julienschmidt/httprouter"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/shots-fired/shots-store/apis"
	"github.com/shots-fired/shots-store/mocks"
	"github.com/shots-fired/shots-store/streamers"
)

func TestAPIs(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "APIs Suite")
}

var _ = Describe("Streamers", func() {
	var (
		ctrl      *gomock.Controller
		mockStore *mocks.MockStore
		router    *httprouter.Router
	)

	BeforeEach(func() {
		ctrl = gomock.NewController(GinkgoT())
		mockStore = mocks.NewMockStore(ctrl)
		router = apis.NewRouter(streamers.NewEngine(mockStore))
	})

	AfterEach(func() {
		ctrl.Finish()
	})

	Describe("Streamers APIs", func() {
		Describe("GET /streamers/{id}", func() {
			It("should return streamer if store returns a valid streamer", func() {
				mockStore.EXPECT().Get(gomock.Any(), gomock.Any()).Return("{\"name\":\"something\"}", nil)
				req, _ := http.NewRequest("GET", "/streamers/1", nil)
				res := executeRequest(req, router)

				Expect(res.Code).To(Equal(http.StatusOK))
				Expect(res.Body.String()).To(Equal("{\"name\":\"something\",\"status\":\"\",\"viewers\":0}"))
			})

			It("should error if the store errors", func() {
				mockStore.EXPECT().Get(gomock.Any(), gomock.Any()).Return("", errors.New("error"))
				req, _ := http.NewRequest("GET", "/streamers/1", nil)
				res := executeRequest(req, router)

				Expect(res.Code).To(Equal(http.StatusInternalServerError))
			})
		})

		Describe("GET /streamers", func() {
			It("should return streamers if store returns a valid streamer", func() {
				mockStore.EXPECT().GetAll(gomock.Any()).Return(map[string]string{
					"1": "{\"name\":\"something\"}",
					"2": "{\"name\":\"else\"}",
				}, nil)
				req, _ := http.NewRequest("GET", "/streamers", nil)
				res := executeRequest(req, router)

				Expect(res.Code).To(Equal(http.StatusOK))
				Expect(res.Body.String()).To(Equal("[{\"name\":\"something\",\"status\":\"\",\"viewers\":0},{\"name\":\"else\",\"status\":\"\",\"viewers\":0}]"))
			})

			It("should error if the store errors", func() {
				mockStore.EXPECT().GetAll(gomock.Any()).Return(map[string]string{}, errors.New("error"))
				req, _ := http.NewRequest("GET", "/streamers", nil)
				res := executeRequest(req, router)

				Expect(res.Code).To(Equal(http.StatusInternalServerError))
			})

			It("should return an empty array if store returns no streamers", func() {
				mockStore.EXPECT().GetAll(gomock.Any()).Return(map[string]string{}, nil)
				req, _ := http.NewRequest("GET", "/streamers", nil)
				res := executeRequest(req, router)

				Expect(res.Code).To(Equal(http.StatusOK))
				Expect(res.Body.String()).To(Equal("[]"))
			})
		})

		Describe("DELETE /streamers/{id}", func() {
			It("should succeed if the store succeeds", func() {
				mockStore.EXPECT().Delete(gomock.Any(), gomock.Any()).Return(nil)
				req, _ := http.NewRequest("DELETE", "/streamers/1", nil)
				res := executeRequest(req, router)

				Expect(res.Code).To(Equal(http.StatusOK))
			})

			It("should error if the store errors", func() {
				mockStore.EXPECT().Delete(gomock.Any(), gomock.Any()).Return(errors.New("error"))
				req, _ := http.NewRequest("DELETE", "/streamers/1", nil)
				res := executeRequest(req, router)

				Expect(res.Code).To(Equal(http.StatusInternalServerError))
			})
		})

		Describe("POST /streamers/{id}", func() {
			It("should succeed if the store succeeds", func() {
				mockStore.EXPECT().Set(gomock.Any(), gomock.Any(), gomock.Any()).Return(nil)
				req, _ := http.NewRequest("POST", "/streamers/1", strings.NewReader("{\"name\":\"something\",\"status\":\"online\",\"viewers\":1}"))
				res := executeRequest(req, router)

				Expect(res.Code).To(Equal(http.StatusOK))
			})

			It("should error if the request body is not proper json", func() {
				req, _ := http.NewRequest("POST", "/streamers/1", strings.NewReader("hi"))
				res := executeRequest(req, router)

				Expect(res.Code).To(Equal(http.StatusBadRequest))
			})

			It("should error if the store errors", func() {
				mockStore.EXPECT().Set(gomock.Any(), gomock.Any(), gomock.Any()).Return(errors.New("error"))
				req, _ := http.NewRequest("POST", "/streamers/1", strings.NewReader("{\"name\":\"something\",\"status\":\"online\",\"viewers\":1}"))
				res := executeRequest(req, router)

				Expect(res.Code).To(Equal(http.StatusInternalServerError))
			})
		})

		Describe("PUT /streamers/{id}", func() {
			It("should succeed if the store succeeds", func() {
				mockStore.EXPECT().Set(gomock.Any(), gomock.Any(), gomock.Any()).Return(nil)
				req, _ := http.NewRequest("PUT", "/streamers/1", strings.NewReader("{\"name\":\"something\",\"status\":\"online\",\"viewers\":1}"))
				res := executeRequest(req, router)

				Expect(res.Code).To(Equal(http.StatusOK))
			})

			It("should error if the request body is not proper json", func() {
				req, _ := http.NewRequest("PUT", "/streamers/1", strings.NewReader("hi"))
				res := executeRequest(req, router)

				Expect(res.Code).To(Equal(http.StatusBadRequest))
			})

			It("should error if the store errors", func() {
				mockStore.EXPECT().Set(gomock.Any(), gomock.Any(), gomock.Any()).Return(errors.New("error"))
				req, _ := http.NewRequest("PUT", "/streamers/1", strings.NewReader("{\"name\":\"something\",\"status\":\"online\",\"viewers\":1}"))
				res := executeRequest(req, router)

				Expect(res.Code).To(Equal(http.StatusInternalServerError))
			})
		})
	})

	Describe("NewRouter", func() {
		It("should not be nil", func() {
			Expect(apis.NewRouter(streamers.NewEngine(mockStore))).ToNot(BeNil())
		})
	})

	Describe("NewServer", func() {
		It("should not be nil", func() {
			Expect(apis.NewServer(httprouter.New())).ToNot(BeNil())
		})
	})
})

func executeRequest(req *http.Request, router *httprouter.Router) *httptest.ResponseRecorder {
	rr := httptest.NewRecorder()
	router.ServeHTTP(rr, req)

	return rr
}
