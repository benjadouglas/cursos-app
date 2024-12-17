<script lang="ts">
    import * as Card from "$lib/components/ui/card/index.js";
    import { Input } from "$lib/components/ui/input";
    import { enhance } from "$app/forms";
    import { Button } from "$lib/components/ui/button";
    import { Checkbox } from "$lib/components/ui/checkbox";
    let checked = $state(false);
</script>

<div class="grid justify-center h-[100vh] content-center">
    <Card.Root class="w-[351px] shadow-lg dark:shadow-gray-900">
        <Card.Header>
            <Card.Title>Create Your account</Card.Title>
            <Card.Description>
                Fill the form below to create your account
            </Card.Description>
        </Card.Header>
        <Card.Content class="">
            <form use:enhance method="POST" class="space-y-2">
                <!-- email -->
                <div class="space-y-2">
                    <p class="data-[fs-error]:text-destructive">email</p>
                    <Input name="email" type="email" />
                </div>
                <!-- username -->
                <div class="space-y-2">
                    <p class="data-[fs-error]:text-destructive">username</p>
                    <Input name="username" type="text" />
                </div>
                <!-- password -->
                <div class="space-y-2">
                    <p class="data-[fs-error]:text-destructive">password</p>
                    <Input name="password" type="password" />
                </div>
                <!-- TODO: Agregar funcionalidad al checkbox -->
                <div
                    class="flex flex-row items-center space-x-2 space-y-2 text-sm font-medium leading-none peer-disabled:cursor-not-allowed peer-disabled:opacity-70"
                >
                    <Checkbox bind:checked />
                    <input type="hidden" name="check" value={checked} />
                    <p class="font-normal text-center mt-0 h-5">
                        Check if you are a administrator
                    </p>
                </div>
                <!-- submit -->
                <Button type="submit" class="mt-2">Submit</Button>
            </form>
        </Card.Content>
    </Card.Root>
</div>
