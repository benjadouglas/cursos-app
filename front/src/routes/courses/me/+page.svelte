<script lang="ts">
    import Curso from "$lib/components/ui/course/course.svelte";
    import type { Course } from "$lib/types";
    import type { PageData } from "./$types";

    let { data }: { data: PageData } = $props();
    let courses: Course[] = data.courses;
</script>

<div
    class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-4 w-full justify-items-center mt-5"
>
    {#each courses as course}
        <Curso
            title={course.Nombre}
            price={course.Precio}
            capacity={course.Capacidad}
            id={course.Id}
            maximo={course.Maximo}
        />
    {/each}
</div>
