<script lang="ts">
    import { Button } from "$lib/components/ui/button";
    import * as Card from "$lib/components/ui/card";

    let { title, price, capacity, maximo, id } = $props();
</script>

<Card.Root class="w-[350px]">
    <Card.Header>
        <Card.Title class="text-xl">{title}</Card.Title>
    </Card.Header>
    <Card.Content>
        <div class="space-y-4">
            <div class="flex justify-between items-center">
                <span class="text-muted-foreground">Price</span>
                <span class="font-semibold">${price}</span>
            </div>
            <div class="flex justify-between items-center">
                <span class="text-muted-foreground">Enrolled</span>
                <span class="font-semibold">{capacity}/{maximo}</span>
            </div>
        </div>
    </Card.Content>
    <Card.Footer class="flex justify-center">
        <Button class="w-full" href="/courses/admin/{id}">Edit</Button>
    </Card.Footer>
</Card.Root>
