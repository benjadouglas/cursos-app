import type { Actions } from "./$types";
import { fail, redirect } from "@sveltejs/kit";

export const actions = {
    default: async ({ request, cookies }) => {
        const form = await request.formData();
        const email = form.get("email") as string;
        const password = form.get("password") as string;
        const response = await fetch("http://localhost:8085/api/login", {
            method: "POST",
            headers: {
                "Content-Type": "application/json",
            },
            mode: "cors",
            body: JSON.stringify({
                email: email,
                password: password,
            }),
        });
        if (!response.ok) {
            return fail(500);
        }
        const data = await response.json();
        cookies.set("session", data.token, {
            path: "/",
            httpOnly: true,
            maxAge: 60 * 60 * 24 * 4, // 4 days
        });
        cookies.set("userId", data.user.ID, {
            path: "/",
            httpOnly: true,
            maxAge: 60 * 60 * 24 * 4, // 4 days
        });
        redirect(301, "/courses");
    },
} satisfies Actions;
