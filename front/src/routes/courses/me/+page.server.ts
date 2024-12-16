import type { PageServerLoad } from "./$types";

// export const load: PageServerLoad = async ({ params }) => {
//     const response = await fetch(
//         `http://localhost:8085/enrollments/${params.slug}`,
//         {
//             method: "GET",
//             headers: {
//                 "Content-Type": "application/json",
//             },
//         },
//     );
//     if (!response.ok) {
//         return fail(500);
//     }
//     const course: Course = await response.json();
//     return {
//         course: course,
//     };
// };
