<script lang="ts">
    import { Input } from "$lib/components/ui/input";
    import type { PageData, ActionData } from "./$types";
    import { Button } from "$lib/components/ui/button";
    import Course from "$lib/components/ui/course/course.svelte";
    import { enhance } from "$app/forms";

    let { data, form }: { data: PageData; form: ActionData } = $props();

    let courses = $derived(form?.courses ?? data.courses);
</script>

<div class="min-h-screen bg-background">
    <main class="container mx-auto px-4">
        <div
            class="flex flex-col items-center justify-center min-h-[80vh] space-y-8"
        >
            <div class="text-center space-y-4 mt-40">
                <h1
                    class="text-4xl font-bold tracking-tighter sm:text-5xl md:text-7xl"
                >
                    Find What You're Looking For
                </h1>
            </div>

            <!-- Search Box -->
            <div class="w-full max-w-2xl mx-auto">
                <form
                    method="POST"
                    use:enhance
                    action="?/update"
                    class="flex gap-2"
                >
                    <Input
                        type="text"
                        name="search"
                        placeholder="Search anything..."
                        class="h-12 text-lg"
                    />
                    <Button type="submit" class="h-12 px-8">Search</Button>
                </form>
            </div>
            <div
                class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-4 w-full justify-items-center"
            >
                {#each courses as item}
                    <Course
                        title={item.Nombre}
                        price={item.Precio}
                        capacity={item.Capacidad}
                        maximo={item.Maximo}
                        id={item.Id}
                    />
                {/each}
            </div>
        </div>
    </main>
</div>
