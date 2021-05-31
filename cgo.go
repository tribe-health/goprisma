package goprisma

/*
#cgo LDFLAGS: -pthread
#cgo darwin,!arm64 LDFLAGS: -L"${SRCDIR}/lib/darwin" -lquery_engine_c_api
#cgo darwin,arm64 LDFLAGS: -L"${SRCDIR}/lib/darwin-aarch64" -lquery_engine_c_api
#cgo linux LDFLAGS: -L"${SRCDIR}/lib/linux" -lquery_engine_c_api
#cgo windows LDFLAGS: -L"${SRCDIR}/lib/windows" -lquery_engine_c_api
*/
import "C"
