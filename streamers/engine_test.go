package streamers_test

import (
	"errors"
	"testing"

	"github.com/golang/mock/gomock"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/shots-fired/shots-store/mocks"
	"github.com/shots-fired/shots-store/streamers"
)

func TestStreamers(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Streamers Suite")
}

var _ = Describe("Streamers", func() {
	var (
		ctrl      *gomock.Controller
		mockStore *mocks.MockStore
		engine    streamers.Engine
	)

	BeforeEach(func() {
		ctrl = gomock.NewController(GinkgoT())
		mockStore = mocks.NewMockStore(ctrl)
		engine = streamers.NewEngine(mockStore)
	})

	AfterEach(func() {
		ctrl.Finish()
	})

	Describe("GetStreamer", func() {
		It("should return a streamer if the store returns data", func() {
			mockStore.EXPECT().Get(gomock.Any(), gomock.Any()).Return("{\"name\":\"something\"}", nil)
			res, err := engine.GetStreamer("field")

			Expect(err).To(BeNil())
			Expect(res).To(Equal(streamers.Streamer{
				Name: "something",
			}))
		})

		It("should error if the store errors", func() {
			mockStore.EXPECT().Get(gomock.Any(), gomock.Any()).Return("{\"name\":\"something\"}", errors.New("error"))
			res, err := engine.GetStreamer("field")

			Expect(err).To(Equal(errors.New("error")))
			Expect(res).To(Equal(streamers.Streamer{}))
		})

		It("should error if the store returns content that cannot be marshaled into a Streamer struct", func() {
			mockStore.EXPECT().Get(gomock.Any(), gomock.Any()).Return("hi", nil)
			res, err := engine.GetStreamer("field")

			Expect(err).ToNot(BeNil())
			Expect(res).To(Equal(streamers.Streamer{}))
		})
	})

	Describe("GetAllStreamers", func() {
		It("should return a streamer if the store returns data", func() {
			mockStore.EXPECT().GetAll(gomock.Any()).Return(map[string]string{
				"1": "{\"name\":\"something\"}",
				"2": "{\"name\":\"else\"}",
			}, nil)
			res, err := engine.GetAllStreamers()

			Expect(err).To(BeNil())
			Expect(res).To(Equal(streamers.Streamers{
				{Name: "something", Status: "", Viewers: 0},
				{Name: "else", Status: "", Viewers: 0},
			}))
		})

		It("should error if the store errors", func() {
			mockStore.EXPECT().GetAll(gomock.Any()).Return(map[string]string{
				"1": "{\"name\":\"something\"}",
				"2": "{\"name\":\"else\"}",
			}, errors.New("error"))
			res, err := engine.GetAllStreamers()

			Expect(err).To(Equal(errors.New("error")))
			Expect(res).To(Equal(streamers.Streamers{}))
		})

		It("should error if the store returns content that cannot be marshaled into a Streamer struct", func() {
			mockStore.EXPECT().GetAll(gomock.Any()).Return(map[string]string{"1": "hi"}, nil)
			res, err := engine.GetAllStreamers()

			Expect(err).ToNot(BeNil())
			Expect(res).To(Equal(streamers.Streamers{}))
		})
	})

	Describe("SetStreamer", func() {
		It("should not error if the store's set does not error", func() {
			mockStore.EXPECT().Set(gomock.Any(), gomock.Any(), gomock.Any()).Return(nil)

			Expect(engine.SetStreamer("field", streamers.Streamer{Name: "something"})).To(BeNil())
		})

		It("should error if the store's set errors", func() {
			mockStore.EXPECT().Set(gomock.Any(), gomock.Any(), gomock.Any()).Return(errors.New("error"))

			Expect(engine.SetStreamer("field", streamers.Streamer{Name: "something"})).To(Equal(errors.New("error")))
		})
	})

	Describe("DeleteStreamer", func() {
		It("should not error if store's delete does not error", func() {
			mockStore.EXPECT().Delete(gomock.Any(), gomock.Any()).Return(nil)

			Expect(engine.DeleteStreamer("field")).To(BeNil())
		})

		It("should error on if store's delete errors", func() {
			mockStore.EXPECT().Delete(gomock.Any(), gomock.Any()).Return(errors.New("error"))

			Expect(engine.DeleteStreamer("field")).To(Equal(errors.New("error")))
		})
	})
})
