<script lang="ts">
    import { Input } from "$lib/components/ui/input";
    import type { PageData } from "./$types";
    import { Button } from "$lib/components/ui/button";
    import { Search } from "lucide-svelte";
    import Course from "$lib/components/ui/course/+page.svelte";

    let { data }: { data: PageData } = $props();
    console.log(data);
    let searchQuery = $state("");
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
                <div class="flex gap-2">
                    <div class="relative flex-1">
                        <Search
                            class="absolute left-3 top-1/2 h-4 w-4 -translate-y-1/2 text-muted-foreground"
                        />
                        <Input
                            type="text"
                            placeholder="Search anything..."
                            class="pl-10 h-12 text-lg"
                            bind:value={searchQuery}
                        />
                    </div>
                    <Button class="h-12 px-8">Search</Button>
                </div>
            </div>
            <div
                class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-4 w-full justify-items-center"
            >
                {#each data.courses as item}
                    <Course
                        title={item.Nombre}
                        price={item.Precio}
                        capacity={item.Capacidad}
                    />
                {/each}
            </div>
        </div>
    </main>
</div>
