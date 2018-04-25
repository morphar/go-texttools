package texttools

import (
	"io/ioutil"
	"testing"
)

type sample struct {
	in, out string
}

func TestShortenWithAppend(t *testing.T) {
	samples := []sample{
		{"sampleText", "sampleText"},                                                                  // Short w/no spaces
		{"sample text", "sample text"},                                                                // Short w/spaces
		{"sampleTextThatIsTooLongWithNoSpaces", "sampleTextThatIsT..."},                               // Long w/no spaces
		{"sampleTextThatIs.Something.", "sampleTextThatIs..."},                                        // Long w/no space - remove the dot
		{"sample text that? I don't know...", "sample text that?.."},                                  // Long w/space - keep the question mark
		{"sample text that.  \r\n \r\n\r\n   \r\n\r\n   Has some punctuation", "sample text that..."}, // Long w/space - remove the dot + space and line stuff
		{"sample text that . Has a dot in the middle", "sample text that..."},                         // Long w/space - remove the dot
		{"sample text\r\n \r\n\r\n   \r\n\r\n  that has lots of space like chars", "sample text that..."},
		{" \r\n \r\n\r\n   \r\n\r\n   ", ""},
	}

	for _, sample := range samples {
		if out := Shorten(sample.in, 20, "..."); out != sample.out {
			t.Errorf("got %q from %q, expected %q", out, sample.in, sample.out)
		}
	}
}

func TestShortenWithoutAppend(t *testing.T) {
	samples := []sample{
		{"sampleText", "sampleText"},                                                                // Short w/no spaces
		{"sample text", "sample text"},                                                              // Short w/spaces
		{"sampleTextThatIsTooLongWithNoSpaces", "sampleTextThatIsTooL"},                             // Long w/no spaces
		{"sampleTextThatIsToo.Something.", "sampleTextThatIsToo."},                                  // Long w/no space - remove the dot
		{"sample text that? I don't know...", "sample text that? I"},                                // Long w/space - keep the question mark
		{"sample text that.  \r\n \r\n\r\n   \r\n\r\n   Has some punctuation", "sample text that."}, // Long w/space - remove the dot + space and line stuff
		{"sample text that . Has a dot in the middle", "sample text that."},
		{"sample text\r\n \r\n\r\n   \r\n\r\n  that has lots of space like chars", "sample text that has"},
		{" \r\n \r\n\r\n   \r\n\r\n   ", ""},
	}

	for _, sample := range samples {
		if out := Shorten(sample.in, 20, ""); out != sample.out {
			t.Errorf("got %q from %q, expected %q", out, sample.in, sample.out)
		}
	}
}

func TestSlug(t *testing.T) {
	samples := []sample{
		{"sample text", "sample-text"},
		{"sample-text", "sample-text"},
		{"sample_text", "sample-text"},
		{"sample___text", "sample-text"},
		{"sampleText", "sample-text"},
		{"inviteYourCustomersAddInvites", "invite-your-customers-add-invites"},
		{"sample 2 Text", "sample-2-text"},
		{"   sample   2    Text   ", "sample-2-text"},
		{"   $#$sample   2    Text   ", "sample-2-text"},
		{"SAMPLE 2 TEXT", "sample-2-text"},
		{"___$$Base64Encode", "base64-encode"},
		{"---$$Base64-_-_-Encode", "base64-encode"},
		{"FOO:BAR$BAZ", "foo-bar-baz"},
		{"FOO#BAR#BAZ", "foo-bar-baz"},
		{"something.com", "something-com"},
		{"$something%", "something"},
		{"something.com", "something-com"},
		{"•¶§ƒ˚foo˙∆˚¬", "foo"},
		{"æøåäò", "aeoaao"},
	}

	for _, sample := range samples {
		if out := Slug(sample.in); out != sample.out {
			t.Errorf("got %q from %q, expected %q", out, sample.in, sample.out)
		}
	}
}

func TestUnCase(t *testing.T) {
	samples := []sample{
		{"sample text", "Sample text"},
		{"sample-text", "Sample text"},
		{"sample_text", "Sample text"},
		{"sample___text", "Sample text"},
		{"sampleText", "Sample text"},
		{"inviteYourCustomersAddInvites", "Invite your customers add invites"},
		{"sample 2 Text", "Sample 2 text"},
		{"   sample   2    Text   ", "Sample 2 text"},
		{"   $#$sample   2    Text   ", "Sample 2 text"},
		{"SAMPLE 2 TEXT", "Sample 2 text"},
		{"___$$Base64Encode", "Base64 encode"},
		{"---$$Base64-_-_-Encode", "Base64 encode"},
		{"FOO:BAR$BAZ", "Foo bar baz"},
		{"FOO#BAR#BAZ", "Foo bar baz"},
		{"something.com", "Something com"},
		{"$something%", "Something"},
		{"something.com", "Something com"},
		{"•¶§ƒ˚foo˙∆˚¬", "Foo"},
	}

	for _, sample := range samples {
		if out := UnCase(sample.in); out != sample.out {
			t.Errorf("got %q from %q, expected %q", out, sample.in, sample.out)
		}
	}
}

func TestSnakeCase(t *testing.T) {
	samples := []sample{
		{"sample text", "sample_text"},
		{"sample-text", "sample_text"},
		{"sample_text", "sample_text"},
		{"sample___text", "sample_text"},
		{"sampleText", "sample_text"},
		{"inviteYourCustomersAddInvites", "invite_your_customers_add_invites"},
		{"sample 2 Text", "sample_2_text"},
		{"   sample   2    Text   ", "sample_2_text"},
		{"   $#$sample   2    Text   ", "sample_2_text"},
		{"SAMPLE 2 TEXT", "sample_2_text"},
		{"___$$Base64Encode", "base64_encode"},
		{"---$$Base64-_-_-Encode", "base64_encode"},
		{"FOO:BAR$BAZ", "foo_bar_baz"},
		{"FOO#BAR#BAZ", "foo_bar_baz"},
		{"something.com", "something_com"},
		{"$something%", "something"},
		{"something.com", "something_com"},
		{"•¶§ƒ˚foo˙∆˚¬", "foo"},
	}

	for _, sample := range samples {
		if out := SnakeCase(sample.in); out != sample.out {
			t.Errorf("got %q from %q, expected %q", out, sample.in, sample.out)
		}
	}
}

func TestKebabCase(t *testing.T) {
	samples := []sample{
		{"sample text", "sample-text"},
		{"sample-text", "sample-text"},
		{"sample_text", "sample-text"},
		{"sample___text", "sample-text"},
		{"sampleText", "sample-text"},
		{"inviteYourCustomersAddInvites", "invite-your-customers-add-invites"},
		{"sample 2 Text", "sample-2-text"},
		{"   sample   2    Text   ", "sample-2-text"},
		{"   $#$sample   2    Text   ", "sample-2-text"},
		{"SAMPLE 2 TEXT", "sample-2-text"},
		{"___$$Base64Encode", "base64-encode"},
		{"---$$Base64-_-_-Encode", "base64-encode"},
		{"FOO:BAR$BAZ", "foo-bar-baz"},
		{"FOO#BAR#BAZ", "foo-bar-baz"},
		{"something.com", "something-com"},
		{"$something%", "something"},
		{"something.com", "something-com"},
		{"•¶§ƒ˚foo˙∆˚¬", "foo"},
	}

	for _, sample := range samples {
		if out := KebabCase(sample.in); out != sample.out {
			t.Errorf("got %q from %q, expected %q", out, sample.in, sample.out)
		}
	}
}

func TestCamelCase(t *testing.T) {
	samples := []sample{
		{"sample text", "sampleText"},
		{"sample-text", "sampleText"},
		{"sample_text", "sampleText"},
		{"sample___text", "sampleText"},
		{"sampleText", "sampleText"},
		{"inviteYourCustomersAddInvites", "inviteYourCustomersAddInvites"},
		{"sample 2 Text", "sample2Text"},
		{"   sample   2    Text   ", "sample2Text"},
		{"   $#$sample   2    Text   ", "sample2Text"},
		{"SAMPLE 2 TEXT", "sample2Text"},
		{"___$$Base64Encode", "base64Encode"},
		{"---$$Base64-_-_-Encode", "base64Encode"},
		{"FOO:BAR$BAZ", "fooBarBaz"},
		{"FOO#BAR#BAZ", "fooBarBaz"},
		{"something.com", "somethingCom"},
		{"$something%", "something"},
		{"something.com", "somethingCom"},
		{"•¶§ƒ˚foo˙∆˚¬", "foo"},
	}

	for _, sample := range samples {
		if out := CamelCase(sample.in); out != sample.out {
			t.Errorf("got %q from %q, expected %q", out, sample.in, sample.out)
		}
	}
}

func TestPascalCase(t *testing.T) {
	samples := []sample{
		{"sample text", "SampleText"},
		{"sample-text", "SampleText"},
		{"sample_text", "SampleText"},
		{"sample___text", "SampleText"},
		{"sampleText", "SampleText"},
		{"inviteYourCustomersAddInvites", "InviteYourCustomersAddInvites"},
		{"sample 2 Text", "Sample2Text"},
		{"   sample   2    Text   ", "Sample2Text"},
		{"   $#$sample   2    Text   ", "Sample2Text"},
		{"SAMPLE 2 TEXT", "Sample2Text"},
		{"___$$Base64Encode", "Base64Encode"},
		{"---$$Base64-_-_-Encode", "Base64Encode"},
		{"FOO:BAR$BAZ", "FooBarBaz"},
		{"FOO#BAR#BAZ", "FooBarBaz"},
		{"something.com", "SomethingCom"},
		{"$something%", "Something"},
		{"something.com", "SomethingCom"},
		{"•¶§ƒ˚foo˙∆˚¬", "Foo"},
	}

	for _, sample := range samples {
		if out := PascalCase(sample.in); out != sample.out {
			t.Errorf("got %q from %q, expected %q", out, sample.in, sample.out)
		}
	}
}

func TestSanitizeText(t *testing.T) {
	samples := []sample{
		{"Text with a half char: ½", "Text with a half char: 1/2"},
		{"Text <b>with</b> a half char: ½", "Text with a half char: 1/2"},
		{"Text <b>with</b> a half char: <a href='_blank'>½</a>", "Text with a half char: 1/2"},
		{`<a href="/shop/cms-9.html">Reparation</a>`, "Reparation"},
	}

	for _, sample := range samples {
		if out := SanitizeText(sample.in); out != sample.out {
			t.Errorf("got %q from %q, expected %q", out, sample.in, sample.out)
		}
	}
}

func TestCP1258ToUTF8(t *testing.T) {
	// The test chars are kept in a Window Latin 1 encoding file,
	// to ensure they're not corrupted by an editor ;)
	f, _ := ioutil.ReadFile("cp1258.txt")
	utf8Str := CP1258ToUTF8(f)
	expected := "€éæøå\r\n"
	if utf8Str != expected {
		t.Errorf("got %q from %q, expected %q", utf8Str, f, expected)
	}
}

func TestRandomStringLength(t *testing.T) {
	if l := len(RandomString(2)); l != 2 {
		t.Errorf("got length: %d, expected 2", l)
	}
	if l := len(RandomString(10)); l != 10 {
		t.Errorf("got length: %d, expected 10", l)
	}
	if l := len(RandomString(20)); l != 20 {
		t.Errorf("got length: %d, expected 20", l)
	}
}

func TestRandomString(t *testing.T) {
	// A simple "collision" test over 3 chars and 2000 runs.
	// Statistically there should be no collsions.
	// With 5 letters and 62 possible letters, there is 916,132,832 possibilities
	var strs []string
	collisions := 0
	for i := 0; i < 2000; i++ {
		str := RandomString(5)
		if StringInSlice(str, strs) {
			collisions++
		}
		strs = append(strs, str)
	}
	// We'll allow for 1 collision, to ensure most tests will succeed
	if collisions > 1 {
		t.Errorf("got %d collisions, expected 0", collisions)
	}
}

func TestRandomStringChars(t *testing.T) {
	// Ensure that all chars in charBytes are used over time.
	// This ensures the random max int is correct.
	var chars []string
	possibleCharsLen := len(possibleChars)

	for i := 0; i < 2000; i++ {
		str := RandomString(10)
		strChars := []byte(str)
		charsLen := len(strChars)
		for n := 0; n < charsLen; n++ {
			curCharStr := string(strChars[n])
			if !StringInSlice(curCharStr, chars) {
				chars = append(chars, curCharStr)
				if len(chars) == possibleCharsLen {
					return
				}
			}
		}
	}

	t.Errorf("got %d chars, expected: %d", len(chars), possibleCharsLen)
}

func BenchmarkShorten(t *testing.B) {
	for i := 0; i < t.N; i++ {
		Shorten("some sample text here_noething:too $amazing.", 20, "...")
	}
}

func BenchmarkSnakeCase(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = SnakeCase("some sample text here_noething:too$amazing")
	}
}

func BenchmarkKebabCase(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = SnakeCase("some sample text here_noething:too$amazing")
	}
}

func BenchmarkCamelCase(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = CamelCase("some sample text here_noething:too$amazing")
	}
}

func BenchmarkPascalCase(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = PascalCase("some sample text here_noething:too$amazing")
	}
}

func BenchmarkRandomString2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = RandomString(2)
	}
}

func BenchmarkRandomString10(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = RandomString(10)
	}
}

func BenchmarkRandomString20(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = RandomString(20)
	}
}
