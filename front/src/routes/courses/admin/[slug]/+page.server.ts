import { fail, type Actions } from "@sveltejs/kit";
import type { PageServerLoad } from "./$types";
import type { Course } from "$lib/types";

export const load: PageServerLoad = async ({ params }) => {
    const response = await fetch(
        `http://localhost:8084/cursos/${params.slug}`,
        {
            method: "GET",
            headers: {
                "Content-Type": "application/json",
            },
        },
    );
    if (!response.ok) {
        return fail(500);
    }
    const course: Course = await response.json();
    return {
        course: course,
    };
};

export const actions = {
    default: async ({ request, cookies, params }) => {
        const form = await request.formData();
        const course_name = form.get("nombre");
        const course_price = parseInt(form.get("precio") as string);
        const user_id = cookies.get("userId");
        const course_id = params.slug;

        const response = await fetch(
            `http://localhost:8084/cursos/${course_id}`,
            {
                method: "PUT",
                headers: {
                    "Content-Type": "application/json",
                    Authorization: `Bearer ${cookies.get("session")}`,
                },
                body: JSON.stringify({
                    Nombre: course_name,
                    Precio: course_price,
                }),
            },
        );
        const data = await response.json();
        if (!response.ok) {
            return fail(500);
        }
        return {
            status: "success",
        };
    },
} satisfies Actions;
