<script lang="ts">
    import * as Card from "$lib/components/ui/card/index.js";
    import { Button } from "$lib/components/ui/button/index.js";
</script>

<div class="grid justify-center h-[100vh] content-center">
    <Card.Root class=" shadow-md dark:shadow-gray-900">
        <Card.Header>
            <Card.Title>Coursera</Card.Title>
            <Card.Description>Aplicacion de cursos</Card.Description>
        </Card.Header>
        <Card.Content class="flex flex-col w-64 gap-y-2">
            <Button href="/login">Login</Button>
            <Button href="/signup" variant="outline">Create account</Button>
        </Card.Content>
    </Card.Root>
</div>
