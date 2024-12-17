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
        <form use:enhance method="POST">
            <Card.Header>
                <Card.Title class="text-2xl font-bold text-center">
                    <input
                        type="text"
                        name="nombre"
                        value={course?.Nombre}
                        class="w-full text-center border rounded-lg p-2"
                    />
                </Card.Title>
            </Card.Header>

            <Card.Content class="space-y-4">
                <div
                    class="flex items-center justify-between p-4 bg-gray-50 rounded-lg"
                >
                    <span class="text-gray-600 font-medium">Price</span>
                    <input
                        type="number"
                        step="0.01"
                        name="precio"
                        value={course?.Precio}
                        class="text-lg font-bold text-right border rounded-lg p-2 w-32"
                    />
                </div>

                <div
                    class="flex items-center justify-between p-4 bg-gray-50 rounded-lg"
                >
                    <span class="text-gray-600 font-medium">Capacity</span>
                    <span class="text-lg font-bold">
                        {course?.Capacidad} students
                    </span>
                </div>
            </Card.Content>
            <Card.Footer class="flex justify-center">
                <input type="hidden" value={course?.Id} name="course_id" />
                <button
                    class="bg-primary text-white px-6 py-2 rounded-lg hover:bg-primary/90 transition-colors"
                    type="submit"
                >
                    Submit Edit
                </button>
            </Card.Footer>
        </form>
    </Card.Root>
</div>
