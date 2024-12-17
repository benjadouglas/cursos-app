import { redirect, type Handle } from "@sveltejs/kit";

export const handle: Handle = async ({ event, resolve }) => {
    const { cookies } = event;
    const userId = cookies.get("userId");
    if (event.url.pathname.startsWith("/courses")) {
        if (userId === "" || userId === undefined) {
            throw redirect(303, "/login");
        }
    }
    if (event.url.pathname === "/login" || event.url.pathname === "/signup") {
        if (userId) {
            throw redirect(303, "/protected");
        }
    }

    if (event.url.pathname.startsWith("/courses/admin")) {
        const response = await fetch("http://localhost:8085/api/validate", {
            method: "POST",
            headers: {
                "Content-Type": "application/json",
                Authorization: `Bearer ${cookies.get("session")}`,
            },
        });
        if (!response.ok) {
            throw redirect(303, "/courses");
        }
    }

    const response = await resolve(event);

    return response;
};
