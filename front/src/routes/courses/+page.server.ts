import type { Actions, PageServerLoad } from "./$types";
import { fail, redirect } from "@sveltejs/kit";

export const load: PageServerLoad = async ({ params }) => {
    const response = await fetch(
        "http://localhost:8087/search?q=Nombre:*~&offset=0&limit=10000",
        {
            method: "GET",
            mode: "cors",
            headers: {
                "Content-Type": "application/json",
            },
        },
    );
    const data = await response.json();
    const sortedData = data.sort((a: any, b: any) => a.Capacidad - b.Capacidad);
    return {
        courses: sortedData,
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
            `http://localhost:8087/search?q=Nombre:${search}*~&offset=0&limit=10000`,
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
        console.log(data);
        return {
            courses: data,
        };
    },
} satisfies Actions;
