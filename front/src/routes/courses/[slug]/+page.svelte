<script lang="ts">
    import type { PageData } from "./$types";
    import * as Card from "$lib/components/ui/card";
    import { enhance } from "$app/forms";
    import type { Course } from "$lib/types";

    let { data }: { data: PageData } = $props();
    let course = data.course;
</script>

<div class="container mx-auto px-4 py-8">
    <Card.Root class="w-full max-w-2xl mx-auto">
        <Card.Header>
            <Card.Title class="text-2xl font-bold text-center"
                >{course?.Nombre}</Card.Title
            >
        </Card.Header>

        <Card.Content class="space-y-4">
            <div
                class="flex items-center justify-between p-4 bg-gray-50 rounded-lg"
            >
                <span class="text-gray-600 font-medium">Price</span>
                <span class="text-lg font-bold">${course?.Precio}</span>
            </div>

            <div
                class="flex items-center justify-between p-4 bg-gray-50 rounded-lg"
            >
                <span class="text-gray-600 font-medium">Capacity</span>
                <span class="text-lg font-bold"
                    >{course?.Capacidad} students</span
                >
            </div>
        </Card.Content>
        <Card.Footer class="flex justify-center">
            <form use:enhance method="POST">
                <input type="hidden" value={course?.id} name="course_id" />
                <button
                    class="bg-primary text-white px-6 py-2 rounded-lg hover:bg-primary/90 transition-colors"
                >
                    Enroll Now
                </button>
            </form>
        </Card.Footer>
    </Card.Root>
</div>
