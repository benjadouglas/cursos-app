<script lang="ts">
    import { BookDashed, Plus, Menu } from "lucide-svelte"
    import { Button } from "$lib/components/ui/button";
    import * as DropdownMenu from "$lib/components/ui/dropdown-menu";
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
                <Button class="h-8" href="/courses/admin/create"><Plus strokeWidth={4} size={16} class="mr-2"/> add a course</Button>
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
                                    class="w-full text-sm font-medium hover:text-primary transition-colors"
                                >
                                    Home
                                </a>
                            </DropdownMenu.Item>
                            <DropdownMenu.Item>
                                <a
                                    href="/courses/me"
                                    class="w-full text-sm font-medium hover:text-primary transition-colors"
                                >
                                    MyCourses
                                </a>
                            </DropdownMenu.Item>
                            <DropdownMenu.Item>
                                <a href="/courses/admin/instances" class="w-full text-sm font-medium hover:text-primary transition-colors">
                                    Instances
                                    </a>
                            </DropdownMenu.Item>
                            <DropdownMenu.Separator />
                            <DropdownMenu.Item>
                                <button
                                    class="w-full text-sm font-medium hover:text-red-500 transition-colors text-left"
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
