package auth

import (
	"net/http"
)

func Getadminapikey(r *http.Request) (string, error) {
    cookie,err := r.Cookie("admin")
    if err != nil {
        return "", err
    }
    return cookie.Value, nil
}

func AdminAuthMiddleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        // Check if the "admin" cookie exists
        _, err := r.Cookie("admin")
        if err != nil {
            // If the cookie doesn't exist, redirect to the login form
            http.Redirect(w, r, "/admin/login/form", http.StatusSeeOther)
            return
        }

        // The cookie exists, proceed to the next handler
        next.ServeHTTP(w, r)
    })
}

func RedirectIfLoggedInMiddleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        // Check if the user is logged in (you may adjust the cookie name)
        _, err := r.Cookie("admin")
        if err == nil {
            // If the user is logged in, redirect to the dashboard
            http.Redirect(w, r, "/admin/dashboard", http.StatusSeeOther)
            return
        }

        // The user is not logged in, proceed to the next handler
        next.ServeHTTP(w, r)
    })
}
