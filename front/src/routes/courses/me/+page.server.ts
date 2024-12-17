import { fail } from "@sveltejs/kit";
import type { PageServerLoad } from "./$types";
import type { Course } from "$lib/types";

export const load: PageServerLoad = async ({ params, cookies }) => {
    const response = await fetch(
        `http://localhost:8085/api/enrollments/user/${cookies.get("userId")}`,
        {
            method: "GET",
            headers: {
                "Content-Type": "application/json",
                Authorization: `Bearer ${cookies.get("session")}`,
            },
        },
    );
    if (!response.ok) {
        return fail(500);
    }
    const courses = await response.json();
    return {
        courses: courses,
    };
};
