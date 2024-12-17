<script lang="ts">
    import type { PageData } from "./$types";

    let { data }: { data: PageData } = $props();
</script>

<div class="p-4">
    <h2 class="text-xl font-semibold mb-2">System Status</h2>

    <div class="space-y-1">
        <div class="flex items-center justify-between p-1 bg-gray-100 rounded">
            <span>Search Service</span>
            <span class={data.solr_ok ? "text-green-600" : "text-red-600"}>
                {data.solr_ok ? "✓" : "✗"}
            </span>
        </div>

        <div class="flex items-center justify-between p-1 bg-gray-100 rounded">
            <span>User Service</span>
            <span class={data.user_ok ? "text-green-600" : "text-red-600"}>
                {data.user_ok ? "✓" : "✗"}
            </span>
        </div>

        <div class="flex items-center justify-between p-1 bg-gray-100 rounded">
            <span>Courses Service</span>
            <span class={data.course_ok ? "text-green-600" : "text-red-600"}>
                {data.course_ok ? "✓" : "✗"}
            </span>
        </div>
    </div>
</div>
