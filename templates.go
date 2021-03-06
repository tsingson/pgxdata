// Do not edit. Automatically generated from templates.go.erb
package main

import (
	"encoding/base64"
	"fmt"
	"text/template"
)

func pkPlaceholder(n int) string {
	return fmt.Sprintf("$%d", n+1)
}

func loadTemplates() *template.Template {
	funcMap := template.FuncMap{"pkPlaceholder": pkPlaceholder}
	templates := template.New("base").Funcs(funcMap)

	var decoded []byte
	var err error

	decoded, err = base64.StdEncoding.DecodeString(`ZnVuYyBVcGRhdGV7ey5TdHJ1Y3ROYW1lfX0oZGIgUXVlcnllcnt7cmFuZ2UgLlByaW1hcnlLZXlDb2x1bW5zfX0sCiAge3suVmFyTmFtZX19IHt7LkdvVHlwZX19e3tlbmR9fSwKICByb3cgKnt7LlN0cnVjdE5hbWV9fSwKKSBlcnJvciB7CiAgc2V0cyA6PSBtYWtlKFtdc3RyaW5nLCAwLCB7e2xlbiAuQ29sdW1uc319KQogIGFyZ3MgOj0gcGd4LlF1ZXJ5QXJncyhtYWtlKFtdaW50ZXJmYWNle30sIDAsIHt7bGVuIC5Db2x1bW5zfX0pKQoKe3tyYW5nZSAuQ29sdW1uc319ICBpZiByb3cue3suRmllbGROYW1lfX0uU3RhdHVzICE9IHBndHlwZS5VbmRlZmluZWQgewogICAgc2V0cyA9IGFwcGVuZChzZXRzLCBge3suQ29sdW1uTmFtZX19YCsiPSIrYXJncy5BcHBlbmQoJnJvdy57ey5GaWVsZE5hbWV9fSkpCiAgfQp7e2VuZH19CgogIGlmIGxlbihzZXRzKSA9PSAwIHsKICAgIHJldHVybiBuaWwKICB9CgogIHNxbCA6PSBgdXBkYXRlICJ7ey5UYWJsZU5hbWV9fSIgc2V0IGAgKyBzdHJpbmdzLkpvaW4oc2V0cywgIiwgIikgKyBgIHdoZXJlIGAge3sgcmFuZ2UgJGksICRjb2x1bW4gOj0gLlByaW1hcnlLZXlDb2x1bW5zfX0gKyBge3tpZiAkaX19IGFuZCB7e2VuZH19Int7JGNvbHVtbi5Db2x1bW5OYW1lfX0iPWAgKyBhcmdzLkFwcGVuZCh7eyRjb2x1bW4uVmFyTmFtZX19KXt7ZW5kfX0KCiAgcHNOYW1lIDo9IHByZXBhcmVkTmFtZSgicGd4ZGF0YVVwZGF0ZXt7LlN0cnVjdE5hbWV9fSIsIHNxbCkKCiAgY29tbWFuZFRhZywgZXJyIDo9IHByZXBhcmVFeGVjKGRiLCBwc05hbWUsIHNxbCwgYXJncy4uLikKICBpZiBlcnIgIT0gbmlsIHsKICAgIHJldHVybiBlcnIKICB9CiAgaWYgY29tbWFuZFRhZy5Sb3dzQWZmZWN0ZWQoKSAhPSAxIHsKICAgIHJldHVybiBFcnJOb3RGb3VuZAogIH0KICByZXR1cm4gbmlsCn0K`)
	if err != nil {
		panic("Unable to decode template")
	}

	_, err = templates.New(`update_func`).Parse(string(decoded))
	if err != nil {
		fmt.Println(string(decoded))
		panic("Unable to parse template")
	}

	decoded, err = base64.StdEncoding.DecodeString(`cGFja2FnZSB7ey5Qa2dOYW1lfX0KLy8gVGhpcyBmaWxlIGlzIGF1dG9tYXRpY2FsbHkgZ2VuZXJhdGVkIGJ5IHBneGRhdGEuCgppbXBvcnQgKAogICJzdHJpbmdzIgoKICAiZ2l0aHViLmNvbS9qYWNrYy9wZ3giCiAgImdpdGh1Yi5jb20vamFja2MvcGd4L3BndHlwZSIKKQoKdHlwZSB7ey5TdHJ1Y3ROYW1lfX0gc3RydWN0IHsKe3tyYW5nZSAuQ29sdW1uc319ICB7ey5GaWVsZE5hbWV9fSB7ey5Hb0JveFR5cGV9fQp7e2VuZH19fQoKe3t0ZW1wbGF0ZSAiY291bnRfZnVuYyIgLn19Cnt7dGVtcGxhdGUgInNlbGVjdF9hbGxfZnVuYyIgLn19Cnt7dGVtcGxhdGUgInNlbGVjdF9ieV9wa19mdW5jIiAufX0Ke3t0ZW1wbGF0ZSAiaW5zZXJ0X2Z1bmMiIC59fQp7e3RlbXBsYXRlICJ1cGRhdGVfZnVuYyIgLn19Cnt7dGVtcGxhdGUgImRlbGV0ZV9mdW5jIiAufX0K`)
	if err != nil {
		panic("Unable to decode template")
	}

	_, err = templates.New(`row`).Parse(string(decoded))
	if err != nil {
		fmt.Println(string(decoded))
		panic("Unable to parse template")
	}

	decoded, err = base64.StdEncoding.DecodeString(`ZnVuYyBJbnNlcnR7ey5TdHJ1Y3ROYW1lfX0oZGIgUXVlcnllciwgcm93ICp7ey5TdHJ1Y3ROYW1lfX0pIGVycm9yIHsKICBhcmdzIDo9IHBneC5RdWVyeUFyZ3MobWFrZShbXWludGVyZmFjZXt9LCAwLCB7e2xlbiAuQ29sdW1uc319KSkKCiAgdmFyIGNvbHVtbnMsIHZhbHVlcyBbXXN0cmluZwoKe3tyYW5nZSAuQ29sdW1uc319ICBpZiByb3cue3suRmllbGROYW1lfX0uU3RhdHVzICE9IHBndHlwZS5VbmRlZmluZWQgewogICAgY29sdW1ucyA9IGFwcGVuZChjb2x1bW5zLCBge3suQ29sdW1uTmFtZX19YCkKICAgIHZhbHVlcyA9IGFwcGVuZCh2YWx1ZXMsIGFyZ3MuQXBwZW5kKCZyb3cue3suRmllbGROYW1lfX0pKQogIH0Ke3tlbmR9fQoKICBzcWwgOj0gYGluc2VydCBpbnRvICJ7ey5UYWJsZU5hbWV9fSIoYCArIHN0cmluZ3MuSm9pbihjb2x1bW5zLCAiLCAiKSArIGApCnZhbHVlcyhgICsgc3RyaW5ncy5Kb2luKHZhbHVlcywgIiwiKSArIGApCnJldHVybmluZyB7eyByYW5nZSAkaSwgJGNvbHVtbiA6PSAuUHJpbWFyeUtleUNvbHVtbnN9fXt7aWYgJGl9fSwge3tlbmR9fSJ7eyRjb2x1bW4uQ29sdW1uTmFtZX19Int7ZW5kfX0KICBgCgogIHBzTmFtZSA6PSBwcmVwYXJlZE5hbWUoInBneGRhdGFJbnNlcnR7ey5TdHJ1Y3ROYW1lfX0iLCBzcWwpCgogIHJldHVybiBwcmVwYXJlUXVlcnlSb3coZGIsIHBzTmFtZSwgc3FsLCBhcmdzLi4uKS5TY2FuKHt7IHJhbmdlICRpLCAkY29sdW1uIDo9IC5QcmltYXJ5S2V5Q29sdW1uc319e3tpZiAkaX19LCB7e2VuZH19JnJvdy57eyRjb2x1bW4uRmllbGROYW1lfX17e2VuZH19KQp9Cg==`)
	if err != nil {
		panic("Unable to decode template")
	}

	_, err = templates.New(`insert_func`).Parse(string(decoded))
	if err != nil {
		fmt.Println(string(decoded))
		panic("Unable to parse template")
	}

	decoded, err = base64.StdEncoding.DecodeString(`cGFja2FnZSB7ey5Qa2dOYW1lfX0KLy8gVGhpcyBmaWxlIGlzIGF1dG9tYXRpY2FsbHkgZ2VuZXJhdGVkIGJ5IHBneGRhdGEuCgppbXBvcnQgKAoJImZtdCIKCSJlcnJvcnMiCgkiaGFzaC9mbnYiCgkiaW8iCgoJImdpdGh1Yi5jb20vamFja2MvcGd4IgopCgpjb25zdCBQR1hEQVRBX1ZFUlNJT04gPSAie3suVmVyc2lvbn19IgoKdmFyIEVyck5vdEZvdW5kID0gZXJyb3JzLk5ldygibm90IGZvdW5kIikKCnR5cGUgUXVlcnllciBpbnRlcmZhY2UgewoJUXVlcnkoc3FsIHN0cmluZywgYXJncyAuLi5pbnRlcmZhY2V7fSkgKCpwZ3guUm93cywgZXJyb3IpCglRdWVyeVJvdyhzcWwgc3RyaW5nLCBhcmdzIC4uLmludGVyZmFjZXt9KSAqcGd4LlJvdwoJRXhlYyhzcWwgc3RyaW5nLCBhcmd1bWVudHMgLi4uaW50ZXJmYWNle30pIChwZ3guQ29tbWFuZFRhZywgZXJyb3IpCn0KCnR5cGUgcHJlcGFyZXIgaW50ZXJmYWNlIHsKCVByZXBhcmUobmFtZSwgc3FsIHN0cmluZykgKCpwZ3guUHJlcGFyZWRTdGF0ZW1lbnQsIGVycm9yKQoJRGVhbGxvY2F0ZShuYW1lIHN0cmluZykgZXJyb3IKfQoKZnVuYyBwcmVwYXJlUXVlcnkoZGIgUXVlcnllciwgbmFtZSwgc3FsIHN0cmluZywgYXJncyAuLi5pbnRlcmZhY2V7fSkgKCpwZ3guUm93cywgZXJyb3IpIHsKCWlmIHByZXBhcmVyLCBvayA6PSBkYi4ocHJlcGFyZXIpOyBvayB7CgkJaWYgXywgZXJyIDo9IHByZXBhcmVyLlByZXBhcmUobmFtZSwgc3FsKTsgZXJyICE9IG5pbCB7CgkJCXJldHVybiBuaWwsIGVycgoJCX0KCQlzcWwgPSBuYW1lCgl9CgoJcmV0dXJuIGRiLlF1ZXJ5KHNxbCwgYXJncy4uLikKfQoKZnVuYyBwcmVwYXJlUXVlcnlSb3coZGIgUXVlcnllciwgbmFtZSwgc3FsIHN0cmluZywgYXJncyAuLi5pbnRlcmZhY2V7fSkgKnBneC5Sb3cgewoJaWYgcHJlcGFyZXIsIG9rIDo9IGRiLihwcmVwYXJlcik7IG9rIHsKCQkvLyBRdWVyeVJvdyBkb2Vzbid0IHJldHVybiBhbiBlcnJvciwgdGhlIGVycm9yIGlzIGVuY29kZWQgaW4gdGhlIHBneC5Sb3cuCgkJLy8gU2luY2UgdGhhdCBpcyBwcml2YXRlLCBJZ25vcmUgdGhlIGVycm9yIGZyb20gUHJlcGFyZSBhbmQgcnVuIHRoZSBxdWVyeQoJCS8vIHdpdGhvdXQgdGhlIHByZXBhcmVkIHN0YXRlbWVudC4gSXQgc2hvdWxkIGZhaWwgd2l0aCB0aGUgc2FtZSBlcnJvci4KCQlpZiBfLCBlcnIgOj0gcHJlcGFyZXIuUHJlcGFyZShuYW1lLCBzcWwpOyBlcnIgPT0gbmlsIHsKCQkJc3FsID0gbmFtZQoJCX0KCX0KCXJldHVybiBkYi5RdWVyeVJvdyhzcWwsIGFyZ3MuLi4pCn0KCmZ1bmMgcHJlcGFyZUV4ZWMoZGIgUXVlcnllciwgbmFtZSwgc3FsIHN0cmluZywgYXJncyAuLi5pbnRlcmZhY2V7fSkgKHBneC5Db21tYW5kVGFnLCBlcnJvcikgewoJaWYgcHJlcGFyZXIsIG9rIDo9IGRiLihwcmVwYXJlcik7IG9rIHsKCQlpZiBfLCBlcnIgOj0gcHJlcGFyZXIuUHJlcGFyZShuYW1lLCBzcWwpOyBlcnIgIT0gbmlsIHsKCQkJcmV0dXJuIHBneC5Db21tYW5kVGFnKCIiKSwgZXJyCgkJfQoJCXNxbCA9IG5hbWUKCX0KCglyZXR1cm4gZGIuRXhlYyhzcWwsIGFyZ3MuLi4pCn0KCmZ1bmMgcHJlcGFyZWROYW1lKGJhc2VOYW1lLCBzcWwgc3RyaW5nKSBzdHJpbmcgewoJaCA6PSBmbnYuTmV3MzJhKCkKCWlmIF8sIGVyciA6PSBpby5Xcml0ZVN0cmluZyhoLCBzcWwpOyBlcnIgIT0gbmlsIHsKCQkvLyBoYXNoLkhhc2guV3JpdGUgbmV2ZXIgcmV0dXJucyBhbiBlcnJvciBzbyB0aGlzIGNhbid0IGhhcHBlbgoJICBwYW5pYygiZmFpbGVkIHdyaXRpbmcgdG8gaGFzaCIpCgl9CgoJcmV0dXJuIGZtdC5TcHJpbnRmKCIlcyVkIiwgYmFzZU5hbWUsIGguU3VtMzIoKSkKfQo=`)
	if err != nil {
		panic("Unable to decode template")
	}

	_, err = templates.New(`db`).Parse(string(decoded))
	if err != nil {
		fmt.Println(string(decoded))
		panic("Unable to parse template")
	}

	decoded, err = base64.StdEncoding.DecodeString(`ZnVuYyBEZWxldGV7ey5TdHJ1Y3ROYW1lfX0oZGIgUXVlcnllcnt7cmFuZ2UgLlByaW1hcnlLZXlDb2x1bW5zfX0sCiAge3suVmFyTmFtZX19IHt7LkdvVHlwZX19e3tlbmR9fSwKKSBlcnJvciB7CiAgYXJncyA6PSBwZ3guUXVlcnlBcmdzKG1ha2UoW11pbnRlcmZhY2V7fSwgMCwge3tsZW4gLlByaW1hcnlLZXlDb2x1bW5zfX0pKQoKICBzcWwgOj0gYGRlbGV0ZSBmcm9tICJ7ey5UYWJsZU5hbWV9fSIgd2hlcmUgYCB7eyByYW5nZSAkaSwgJGNvbHVtbiA6PSAuUHJpbWFyeUtleUNvbHVtbnN9fSArIGB7e2lmICRpfX0gYW5kIHt7ZW5kfX0ie3skY29sdW1uLkNvbHVtbk5hbWV9fSI9YCArIGFyZ3MuQXBwZW5kKHt7JGNvbHVtbi5WYXJOYW1lfX0pe3tlbmR9fQoKICBjb21tYW5kVGFnLCBlcnIgOj0gcHJlcGFyZUV4ZWMoZGIsICJwZ3hkYXRhRGVsZXRle3suU3RydWN0TmFtZX19Iiwgc3FsLCBhcmdzLi4uKQogIGlmIGVyciAhPSBuaWwgewogICAgcmV0dXJuIGVycgogIH0KICBpZiBjb21tYW5kVGFnLlJvd3NBZmZlY3RlZCgpICE9IDEgewogICAgcmV0dXJuIEVyck5vdEZvdW5kCiAgfQogIHJldHVybiBuaWwKfQo=`)
	if err != nil {
		panic("Unable to decode template")
	}

	_, err = templates.New(`delete_func`).Parse(string(decoded))
	if err != nil {
		fmt.Println(string(decoded))
		panic("Unable to parse template")
	}

	decoded, err = base64.StdEncoding.DecodeString(`cGFja2FnZSA9ICJ7ey5Qa2dOYW1lfX0iCgojIERhdGFiYXNlIGNvbm5lY3Rpb24gaW5mb3JtYXRpb24gY2FuIGJlIHNwZWNpZmllZCBoZXJlIG9yIGluIFBHKiBlbnZpcm9ubWVudCB2YXJpYWJsZXMKIwojIFtkYXRhYmFzZV0KIyBob3N0ID0gIjEyNy4wLjAuMSIKIyBwb3J0ID0gNTQzMgojIGRhdGFiYXNlID0gIm15YXBwX2RldmVsb3BtZW50IgojIHVzZXIgPSAibXl1c2VyIgojIHBhc3N3b3JkID0gInNlY3JldCIKCltbdGFibGVzXV0KdGFibGVfbmFtZSA9ICJjdXN0b21lciIKIyBzdHJ1Y3RfbmFtZSA9ICJDdXN0b21lciIK`)
	if err != nil {
		panic("Unable to decode template")
	}

	_, err = templates.New(`config`).Parse(string(decoded))
	if err != nil {
		fmt.Println(string(decoded))
		panic("Unable to parse template")
	}

	decoded, err = base64.StdEncoding.DecodeString(`Y29uc3Qgc2VsZWN0e3suU3RydWN0TmFtZX19QnlQS1NRTCA9IGBzZWxlY3R7eyByYW5nZSAkaSwgJGNvbHVtbiA6PSAuQ29sdW1uc319e3tpZiAkaX19LHt7ZW5kfX0KICAie3skY29sdW1uLkNvbHVtbk5hbWV9fSJ7e2VuZH19CmZyb20gInt7LlRhYmxlTmFtZX19Igp3aGVyZSB7eyByYW5nZSAkaSwgJGNvbHVtbiA6PSAuUHJpbWFyeUtleUNvbHVtbnN9fXt7aWYgJGl9fSBhbmQge3tlbmR9fSJ7eyRjb2x1bW4uQ29sdW1uTmFtZX19Ij17e3BrUGxhY2Vob2xkZXIgJGl9fXt7ZW5kfX1gCgpmdW5jIFNlbGVjdHt7LlN0cnVjdE5hbWV9fUJ5UEsoCiAgZGIgUXVlcnllcnt7cmFuZ2UgLlByaW1hcnlLZXlDb2x1bW5zfX0sCiAge3suVmFyTmFtZX19IHt7LkdvVHlwZX19e3tlbmR9fSwKKSAoKnt7LlN0cnVjdE5hbWV9fSwgZXJyb3IpIHsKICB2YXIgcm93IHt7LlN0cnVjdE5hbWV9fQogIGVyciA6PSBwcmVwYXJlUXVlcnlSb3coZGIsICJwZ3hkYXRhU2VsZWN0e3suU3RydWN0TmFtZX19QnlQSyIsIHNlbGVjdHt7LlN0cnVjdE5hbWV9fUJ5UEtTUUx7e3JhbmdlIC5QcmltYXJ5S2V5Q29sdW1uc319LCB7ey5WYXJOYW1lfX17e2VuZH19KS5TY2FuKAp7e3JhbmdlIC5Db2x1bW5zfX0mcm93Lnt7LkZpZWxkTmFtZX19LAogICAge3tlbmR9fSkKICBpZiBlcnIgPT0gcGd4LkVyck5vUm93cyB7CiAgICByZXR1cm4gbmlsLCBFcnJOb3RGb3VuZAogIH0gZWxzZSBpZiBlcnIgIT0gbmlsIHsKICAgIHJldHVybiBuaWwsIGVycgogIH0KCiAgcmV0dXJuICZyb3csIG5pbAp9Cg==`)
	if err != nil {
		panic("Unable to decode template")
	}

	_, err = templates.New(`select_by_pk_func`).Parse(string(decoded))
	if err != nil {
		fmt.Println(string(decoded))
		panic("Unable to parse template")
	}

	decoded, err = base64.StdEncoding.DecodeString(`Y29uc3QgY291bnR7ey5TdHJ1Y3ROYW1lfX1TUUwgPSBgc2VsZWN0IGNvdW50KCopIGZyb20gInt7LlRhYmxlTmFtZX19ImAKCmZ1bmMgQ291bnR7ey5TdHJ1Y3ROYW1lfX0oZGIgUXVlcnllcikgKGludDY0LCBlcnJvcikgewogIHZhciBuIGludDY0CiAgZXJyIDo9IHByZXBhcmVRdWVyeVJvdyhkYiwgInBneGRhdGFDb3VudHt7LlN0cnVjdE5hbWV9fSIsIGNvdW50e3suU3RydWN0TmFtZX19U1FMKS5TY2FuKCZuKQogIHJldHVybiBuLCBlcnIKfQo=`)
	if err != nil {
		panic("Unable to decode template")
	}

	_, err = templates.New(`count_func`).Parse(string(decoded))
	if err != nil {
		fmt.Println(string(decoded))
		panic("Unable to parse template")
	}

	decoded, err = base64.StdEncoding.DecodeString(`Y29uc3QgU2VsZWN0QWxse3suU3RydWN0TmFtZX19U1FMID0gYHNlbGVjdHt7IHJhbmdlICRpLCAkY29sdW1uIDo9IC5Db2x1bW5zfX17e2lmICRpfX0se3tlbmR9fQogICJ7eyRjb2x1bW4uQ29sdW1uTmFtZX19Int7ZW5kfX0KZnJvbSAie3suVGFibGVOYW1lfX0iYAoKZnVuYyBTZWxlY3RBbGx7ey5TdHJ1Y3ROYW1lfX0oZGIgUXVlcnllcikgKFtde3suU3RydWN0TmFtZX19LCBlcnJvcikgewogIHZhciByb3dzIFtde3suU3RydWN0TmFtZX19CgogIGRiUm93cywgZXJyIDo9IHByZXBhcmVRdWVyeShkYiwgInBneGRhdGFTZWxlY3RBbGx7ey5TdHJ1Y3ROYW1lfX0iLCBTZWxlY3RBbGx7ey5TdHJ1Y3ROYW1lfX1TUUwpCiAgaWYgZXJyICE9IG5pbCB7CiAgICByZXR1cm4gbmlsLCBlcnIKICB9CgogIGZvciBkYlJvd3MuTmV4dCgpIHsKICAgIHZhciByb3cge3suU3RydWN0TmFtZX19CiAgICBkYlJvd3MuU2NhbigKe3tyYW5nZSAuQ29sdW1uc319JnJvdy57ey5GaWVsZE5hbWV9fSwKICAgIHt7ZW5kfX0pCiAgICByb3dzID0gYXBwZW5kKHJvd3MsIHJvdykKICB9CgogIGlmIGRiUm93cy5FcnIoKSAhPSBuaWwgewogICAgcmV0dXJuIG5pbCwgZGJSb3dzLkVycigpCiAgfQoKICByZXR1cm4gcm93cywgbmlsCn0K`)
	if err != nil {
		panic("Unable to decode template")
	}

	_, err = templates.New(`select_all_func`).Parse(string(decoded))
	if err != nil {
		fmt.Println(string(decoded))
		panic("Unable to parse template")
	}

	return templates
}
