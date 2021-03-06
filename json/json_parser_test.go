package json_test

import (
	"io/ioutil"
	"os"

	"github.com/cloudfoundry/cli/json"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("ParseJson", func() {
	var filename string
	var tmpFile *os.File

	Context("when everything is proper", func() {
		BeforeEach(func() {
			tmpFile, _ = ioutil.TempFile("", "WONDERFULFILEWHOSENAMEISHARDTOREADBUTCONTAINSVALIDJSON")
			filename = tmpFile.Name()
			ioutil.WriteFile(filename, []byte("[{\"akey\": \"avalue\"}]"), 0644)
		})

		AfterEach(func() {
			tmpFile.Close()
			os.Remove(tmpFile.Name())
		})

		It("converts a json file into an unmarshalled slice of string->string map objects", func() {
			stringMaps, err := json.ParseJSON(filename)
			Expect(err).To(BeNil())
			Expect(stringMaps[0]["akey"]).To(Equal("avalue"))
		})
	})

	Context("when the JSON is invalid", func() {
		BeforeEach(func() {
			tmpFile, _ = ioutil.TempFile("", "TERRIBLEFILECONTAININGINVALIDJSONWHICHMAKESEVERYTHINGTERRIBLEANDSTILLHASANAMETHATSHARDTOREAD")
			filename = tmpFile.Name()
			ioutil.WriteFile(filename, []byte("SCARY NOISES}"), 0644)
		})

		AfterEach(func() {
			tmpFile.Close()
			os.Remove(tmpFile.Name())
		})

		It("tries to convert the json file but fails because it was given something it didn't like", func() {
			_, err := json.ParseJSON(filename)
			Expect(err).ToNot(BeNil())
		})
	})
})
