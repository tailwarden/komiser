import CardSkeleton from '../../../card/CardSkeleton';

function DashboardTopStatsSkeleton() {
  const numberOfSkeletonCardsToRender = Array.from(Array(4).keys());

  return (
    <div
      data-testid="loading"
      className="grid-col grid gap-8 md:grid-cols-2 lg:grid-cols-4"
    >
      {numberOfSkeletonCardsToRender.map(skeleton => (
        <CardSkeleton key={skeleton} />
      ))}
    </div>
  );
}

export default DashboardTopStatsSkeleton;
