<script lang="ts">
	import * as Chart from "$lib/components/ui/chart/index.js";
	import * as Card from "$lib/components/ui/card/index.js";
	import * as Select from "$lib/components/ui/select/index.js";
	import * as ToggleGroup from "$lib/components/ui/toggle-group/index.js";
	import { scaleUtc } from "d3-scale";
	import { Area, AreaChart } from "layerchart";
	import { curveNatural } from "d3-shape";

	const complianceData = [
		{ date: new Date("2024-04-15"), compliance: 75 },
		{ date: new Date("2024-04-28"), compliance: 82 },
		{ date: new Date("2024-05-10"), compliance: 68 },
		{ date: new Date("2024-05-22"), compliance: 79 },
		{ date: new Date("2024-06-05"), compliance: 85 },
		{ date: new Date("2024-06-18"), compliance: 71 },
		{ date: new Date("2024-06-30"), compliance: 67 },
	];

	let timeRange = $state("90d");

	const selectedLabel = $derived.by(() => {
		switch (timeRange) {
			case "90d":
				return "Last 3 months";
			case "30d":
				return "Last 30 days";
			case "7d":
				return "Last 7 days";
			default:
				return "Last 3 months";
		}
	});

	const filteredData = $derived(
		complianceData.filter((item) => {
			const referenceDate = new Date("2024-06-30");
			let daysToSubtract = 90;
			if (timeRange === "30d") {
				daysToSubtract = 30;
			} else if (timeRange === "7d") {
				daysToSubtract = 7;
			}

			referenceDate.setDate(referenceDate.getDate() - daysToSubtract);
			return item.date >= referenceDate;
		})
	);

	const chartConfig = {
		compliance: { label: "Compliance %", color: "var(--chart-1)" },
	} satisfies Chart.ChartConfig;
</script>

<Card.Root class="@container/card">
	<Card.Header>
		<Card.Title>Compliance Over Time</Card.Title>
		<Card.Description>
			<span class="hidden @[540px]/card:block"> Percentage of compliant entries per audit </span>
			<span class="@[540px]/card:hidden">Per audit</span>
		</Card.Description>
		<Card.Action>
			<ToggleGroup.Root
				type="single"
				bind:value={timeRange}
				variant="outline"
				class="hidden *:data-[slot=toggle-group-item]:!px-4 @[767px]/card:flex"
			>
				<ToggleGroup.Item value="90d">Last 3 months</ToggleGroup.Item>
				<ToggleGroup.Item value="30d">Last 30 days</ToggleGroup.Item>
				<ToggleGroup.Item value="7d">Last 7 days</ToggleGroup.Item>
			</ToggleGroup.Root>
			<Select.Root type="single" bind:value={timeRange}>
				<Select.Trigger
					size="sm"
					class="flex w-40 **:data-[slot=select-value]:block **:data-[slot=select-value]:truncate @[767px]/card:hidden"
					aria-label="Select a value"
				>
					<span data-slot="select-value">
						{selectedLabel}
					</span>
				</Select.Trigger>
				<Select.Content class="rounded-xl">
					<Select.Item value="90d" class="rounded-lg">Last 3 months</Select.Item>
					<Select.Item value="30d" class="rounded-lg">Last 30 days</Select.Item>
					<Select.Item value="7d" class="rounded-lg">Last 7 days</Select.Item>
				</Select.Content>
			</Select.Root>
		</Card.Action>
	</Card.Header>
	<Card.Content class="px-2 pt-4 sm:px-6 sm:pt-6">
		<Chart.Container config={chartConfig} class="aspect-auto h-[250px] w-full">
			<AreaChart
				legend
				data={filteredData}
				x="date"
				xScale={scaleUtc()}
				y="compliance"
				series={[
					{
						key: "compliance",
						label: "Compliance %",
						color: chartConfig.compliance.color,
					},
				]}
				props={{
					xAxis: {
						ticks: timeRange === "7d" ? 7 : undefined,
						format: (v) => {
							return v.toLocaleDateString("en-US", {
								month: "short",
								day: "numeric",
							});
						},
					},
					yAxis: { format: () => "" },
				}}
			>
				{#snippet marks({ context })}
					<defs>
						<linearGradient id="fillCompliance" x1="0" y1="0" x2="0" y2="1">
							<stop
								offset="5%"
								stop-color="var(--color-compliance)"
								stop-opacity={1.0}
							/>
							<stop
								offset="95%"
								stop-color="var(--color-compliance)"
								stop-opacity={0.1}
							/>
						</linearGradient>
					</defs>
					{#each context.series.visibleSeries as s (s.key)}
						<Area
							seriesKey={s.key}
							curve={curveNatural}
							fillOpacity={0.4}
							line={{ class: "stroke-1" }}
							motion="tween"
							{...s.props}
							fill="url(#fillCompliance)"
						/>
					{/each}
				{/snippet}
				{#snippet tooltip()}
					<Chart.Tooltip
						labelFormatter={(v: Date) => {
							return v.toLocaleDateString("en-US", {
								month: "short",
								day: "numeric",
							});
						}}
						indicator="line"
					/>
				{/snippet}
			</AreaChart>
		</Chart.Container>
	</Card.Content>
</Card.Root>
