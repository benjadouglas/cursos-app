import type { Actions, PageServerLoad } from "./$types";
import { fail, redirect } from "@sveltejs/kit";

export const load: PageServerLoad = async ({ params }) => {
    const response = await fetch(
        "http://localhost:8087/search?q=Id:*&offset=0&limit=10000",
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
