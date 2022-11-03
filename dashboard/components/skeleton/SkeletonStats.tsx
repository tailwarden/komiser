function SkeletonStats() {
  const cards: number[] = Array.from(Array(3).keys());

  return (
    <div className="grid grid-col md:grid-cols-2 lg:grid-cols-3 gap-8">
      {cards.map(card => (
        <div
          key={card}
          className="flex items-center h-[7.5rem] px-6 text-sm bg-white rounded-lg animate-pulse"
        >
          <div className="w-full flex gap-6">
            <div className="flex-shrink-0 w-10 h-10 rounded-full bg-komiser-200/50"></div>
            <div className="flex flex-col w-full gap-3">
              <div className="w-[36%] h-4 bg-komiser-200/50 rounded-lg"></div>
              <div className="w-[86%] h-4 bg-komiser-200/50 rounded-lg"></div>
            </div>
          </div>
        </div>
      ))}
    </div>
  );
}

export default SkeletonStats;
