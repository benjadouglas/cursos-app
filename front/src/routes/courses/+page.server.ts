import type { Actions, PageServerLoad } from "./$types";
import { fail, redirect } from "@sveltejs/kit";

export const load: PageServerLoad = async ({ params }) => {
    const response = await fetch(
        "http://localhost:8080/search?q=Id:*&offset=0&limit=10000",
        {
            method: "GET",
            mode: "cors",
            headers: {
                "Content-Type": "application/json",
            },
        },
    );
    const data = await response.json();
    return {
        courses: data,
    };
};

export const actions = {
    update: async ({ request, cookies }) => {
        const form = await request.formData();
        let search = form.get("search") as string;
        if (search === "") {
            search = "*";
        }
        const response = await fetch(
            `http://localhost:8080/search?q=Nombre:${search}~&offset=0&limit=10000`,
            {
                method: "GET",
                mode: "cors",
                headers: {
                    "Content-Type": "application/json",
                },
            },
        );

        if (!response.ok) {
            return fail(500);
        }
        const data = await response.json();
        return {
            courses: data,
        };
    },
} satisfies Actions;