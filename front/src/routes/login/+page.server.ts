import type { Actions } from "./$types";
import { fail } from "@sveltejs/kit";

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
        console.log(data);
    },
} satisfies Actions;
