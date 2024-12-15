<script lang="ts">
    // import '../app.css';
    import { invalidateAll } from "$app/navigation";
    import { goto } from "$app/navigation";
    import { Button } from "$lib/components/ui/button";
    let { children } = $props();

    async function handleLogout() {
    }
</script>

<div class="min-h-screen flex flex-col">
    <header class="border-b">
        <div class="container mx-auto px-4">
            <nav class="flex items-center h-16 justify-between">
                <a href="/protected" class="flex items-center space-x-2">
                    <svg
                        xmlns="http://www.w3.org/2000/svg"
                        width="24"
                        height="24"
                        viewBox="0 0 24 24"
                        fill="none"
                        stroke="currentColor"
                        stroke-width="2"
                        stroke-linecap="round"
                        stroke-linejoin="round"
                        class="h-6 w-6"
                    >
                        <path d="M12 2L2 7l10 5 10-5-10-5z" />
                        <path d="M2 17l10 5 10-5" />
                        <path d="M2 12l10 5 10-5" />
                    </svg>
                    <span class="font-bold text-xl">LearnHub</span>
                </a>
                <!-- Navigation Links -->
                <div class="flex items-center space-x-2">
                    <a
                        href="/protected"
                        class="text-sm font-medium hover:text-primary transition-colors"
                    >
                        Home
                    </a>
                    <a
                        href="/protected/mycourses"
                        class="text-sm font-medium hover:text-primary transition-colors"
                    >
                        Courses
                    </a>
                    <button
                        onclick={handleLogout}
                        class="text-sm font-medium"
                    >
                        Logout
                    </button>
                </div>
            </nav>
        </div>
    </header>

    <!-- Page Content -->
    <main class="flex-1">
        {@render children()}
    </main>
</div>
