<script lang="ts">
    import { BookDashed } from "lucide-svelte"
    import * as DropdownMenu from "$lib/components/ui/dropdown-menu";
    import {Menu} from "lucide-svelte";
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
                <DropdownMenu.Root>
                    <DropdownMenu.Trigger>
                        <Button variant="outline" size="icon" class="w-10">
                            <Menu size={20} />
                        </Button>
                    </DropdownMenu.Trigger>
                    <DropdownMenu.Content>
                        <DropdownMenu.Group>
                            <DropdownMenu.Item>
                                <a
                                    href="/courses"
                                    class="w-full text-sm font-normal hover:text-primary transition-colors"
                                >
                                    Home
                                </a>
                            </DropdownMenu.Item>
                            <DropdownMenu.Item>
                                <a
                                    href="/courses/me"
                                    class="w-full text-sm font-normal hover:text-primary transition-colors"
                                >
                                    My Courses
                                </a>
                            </DropdownMenu.Item>
                            <DropdownMenu.Item>
                                <a
                                    href="/courses/admin"
                                    class="w-full text-sm font-normal hover:text-primary transition-colors"
                                >
                                    Admin
                                </a>
                            </DropdownMenu.Item>
                            <DropdownMenu.Separator />
                            <DropdownMenu.Item>
                                <button
                                    class="w-full text-sm font-normal hover:text-red-500 transition-colors text-left"
                                >
                                    Logout
                                </button>
                            </DropdownMenu.Item>
                        </DropdownMenu.Group>
                    </DropdownMenu.Content>
                </DropdownMenu.Root>
            </nav>
        </div>
    </header>

    <!-- Page Content -->
    <main class="flex-1">
        {@render children()}
    </main>
</div>
