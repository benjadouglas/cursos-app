import type { Actions, PageServerLoad } from "./$types";
import { fail, redirect } from "@sveltejs/kit";

export const load: PageServerLoad = async ({ params, cookies }) => {
    let solr_ok = false;
    let user_ok = false;
    let course_ok = false;

    try {
        const solr_response = await fetch(
            "http://localhost:8087/search?q=Id:*&offset=0&limit=10000",
            {
                method: "GET",
                mode: "cors",
                headers: {
                    "Content-Type": "application/json",
                },
            },
        );
        solr_ok = solr_response.ok;
    } catch (error) {
        solr_ok = false;
    }

    try {
        const user_response = await fetch(
            "http://localhost:8085/api/validate",
            {
                method: "POST",
                headers: {
                    "Content-Type": "application/json",
                    Authorization: `Bearer ${cookies.get("session")}`,
                },
            },
        );
        user_ok = user_response.ok;
    } catch (error) {
        user_ok = false;
    }

    try {
        const courses_response = await fetch(
            "http://localhost:8084/cursos/6761eb462b4b4537265d3c5b",
            {
                method: "GET",
                headers: {
                    "Content-Type": "application/json",
                },
            },
        );
        course_ok = courses_response.ok;
    } catch (error) {
        course_ok = false;
    }

    return {
        solr_ok,
        user_ok,
        course_ok,
    };
};
