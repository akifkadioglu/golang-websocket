import router from "./router/index"
export const cryption = {
    parseJwt(token) {
        try {
            var base64Url = token.split('.')[1];
            var base64 = base64Url.replace(/-/g, '+').replace(/_/g, '/');
            var jsonPayload = decodeURIComponent(window.atob(base64).split('').map(function (c) {
                return '%' + ('00' + c.charCodeAt(0).toString(16)).slice(-2);
            }).join(''));

            return JSON.parse(jsonPayload);
        } catch (error) {
            router.replace("/login")
            return
        }
    },

}

export const auth = {
    async check(token) {


        if (!token) {
            localStorage.clear();
            router.replace("/login");
            return;
        }

        var s = await fetch("https://socket-nwnt.onrender.com/check", {
            headers: {
                "Content-Type": "application/json",
                Accept: "application/json",
                authorization: "Bearer " + token,
            },
        });

        if (s.status != 200) {
            localStorage.clear();
            router.replace("/login");
        }
    }
}