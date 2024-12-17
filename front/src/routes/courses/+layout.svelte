<script lang="ts">
    import { BookDashed } from "lucide-svelte"
    import { Button } from "$lib/components/ui/button";
    let { children } = $props();
</script>

<div class="min-h-screen flex flex-col">
    <header class="border-b">
        <div class="container mx-auto px-4">
            <nav class="flex items-center h-16 justify-between">
                <a href="/courses" class="flex items-center space-x-2">
                    <BookDashed/>
                    <span class="font-bold text-xl">LearnHub</span>
                </a>
                <div class="flex items-center space-x-2">
                    <Button
                        variant="link"
                        href="/courses"
                        class="text-sm font-medium hover:text-primary transition-colors"
                    >
                        Home
                    </Button>
                    <Button
                        variant="link"
                        href="/courses/me"
                        class="text-sm font-medium hover:text-primary transition-colors"
                    >
                        MyCourses
                    </Button>
                    <Button
                        variant="link"
                        class="text-sm font-medium hover:text-red-500"

                    >
                        Logout
                    </Button>
                </div>
            </nav>
        </div>
    </header>

    <!-- Page Content -->
    <main class="flex-1">
        {@render children()}
    </main>
</div>
