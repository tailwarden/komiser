type SkeletonStatsProps = {
  NumOfCards: number;
};

function SkeletonStats({ NumOfCards = 3 }) {
  const cards: number[] = Array.from(Array(NumOfCards).keys());

  return (
    <div
      className={`grid-col grid gap-8 md:grid-cols-2 ${
        NumOfCards === 4 ? 'lg:grid-cols-4' : 'lg:grid-cols-3'
      }`}
    >
      {cards.map(card => (
        <div
          key={card}
          className="flex h-[7.5rem] animate-pulse items-center rounded-lg bg-white px-6 text-sm"
        >
          <div className="flex w-full gap-6">
            <div className="h-10 w-10 flex-shrink-0 rounded-xl bg-cyan-200"></div>
            <div className="flex w-full flex-col gap-3">
              <div className="h-4 w-[36%] rounded-lg bg-cyan-200"></div>
              <div className="h-4 w-[86%] rounded-lg bg-cyan-200"></div>
            </div>
          </div>
        </div>
      ))}
    </div>
  );
}

export default SkeletonStats;
