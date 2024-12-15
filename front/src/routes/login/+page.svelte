<script lang="ts">
    import * as Card from "$lib/components/ui/card/index.js";
    import * as Form from "$lib/components/ui/form";
    import { enhance } from "$app/forms";
    import { Input } from "$lib/components/ui/input";
</script>

<div class="grid justify-center h-[100vh] content-center">
    <Card.Root class="w-[351px] shadow-lg dark:shadow-gray-900">
        <Card.Header>
            <Card.Title>Login to your account</Card.Title>
            <Card.Description>
                Enter your email below to login to your account
            </Card.Description>
        </Card.Header>
        <Card.Content class="">
            <form use:enhance method="POST">
                <div class="space-y-2">
                    <p class="data-[fs-error]:text-destructive">Email</p>
                    <Input name="email" type="email" />
                    <p class="data-[fs-error]:text-destructive">Password</p>
                    <Input name="password" type="password" />
                    <Form.Button class="mt-2">Submit</Form.Button>
                </div>
            </form>
        </Card.Content>
    </Card.Root>
</div>
