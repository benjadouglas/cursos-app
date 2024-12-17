import { fail, type Actions } from "@sveltejs/kit";

export const actions = {
    default: async ({ request, cookies }) => {
        const form = await request.formData();
        const course_name = form.get("nombre");
        const course_price = parseInt(form.get("precio") as string);
        const user_id = cookies.get("userId");

        const response = await fetch("http://localhost:8084/cursos", {
            method: "POST",
            headers: {
                "Content-Type": "application/json",
                Authorization: `Bearer ${cookies.get("session")}`,
            },
            body: JSON.stringify({
                Nombre: course_name,
                Precio: course_price,
            }),
        });
        const data = await response.json();
        console.log(data);
        if (!response.ok) {
            return fail(500);
        }
        return {
            status: "success",
        };
    },
} satisfies Actions;
