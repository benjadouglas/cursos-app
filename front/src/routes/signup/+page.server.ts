import type { Actions } from "./$types";
import { fail, redirect } from "@sveltejs/kit";

export const actions = {
    default: async ({ request, cookies }) => {
        const form = await request.formData();
        const username = form.get("username") as string;
        const email = form.get("email") as string;
        const password = form.get("password") as string;
        const response = await fetch("http://localhost:8085/api/users", {
            method: "POST",
            headers: {
                "Content-Type": "application/json",
            },
            body: JSON.stringify({
                nombre: username,
                email: email,
                password: password,
            }),
        });
        if (!response.ok) {
            return fail(500);
        }
        redirect(301, "/login");
    },
} satisfies Actions;
