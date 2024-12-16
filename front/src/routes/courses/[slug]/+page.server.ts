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
    default: async ({ request, cookies }) => {
        const form = await request.formData();
        const user_id = cookies.get("userId");
        const course_id = form.get("course_id");
        console.log({ user_id, course_id });
        const response = await fetch("http://localhost:8085/api/enrollments", {
            method: "POST",
            headers: {
                "Content-Type": "application/json",
                Authorization: `Bearer ${cookies.get("session")}`,
            },
            body: JSON.stringify({
                id: user_id,
                curso_id: course_id,
            }),
        });
        if (!response.ok) {
            return fail(500);
        }
        return {
            status: "success",
        };
    },
} satisfies Actions;
