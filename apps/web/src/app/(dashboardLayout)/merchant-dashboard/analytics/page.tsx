
export default function MerchantAnalyticsPage() {
  return (
    <div className="space-y-6">
      <div>
        <h1 className="text-2xl font-black text-[#404042] tracking-tight">Merchant Analytics</h1>
        <p className="text-xs text-slate-500">Delivery success rate, return stats, & COD trends</p>
      </div>

      <div className="grid grid-cols-2 sm:grid-cols-4 gap-4">
        <div className="bg-white p-5 rounded-3xl border border-slate-200 shadow-sm space-y-1">
          <span className="text-xs font-semibold text-slate-500">Success Rate</span>
          <div className="text-2xl font-black text-emerald-600">96.4%</div>
          <span className="text-[10px] text-slate-400">Past 30 Days</span>
        </div>
        <div className="bg-white p-5 rounded-3xl border border-slate-200 shadow-sm space-y-1">
          <span className="text-xs font-semibold text-slate-500">Total Parcels</span>
          <div className="text-2xl font-black text-[#404042]">482</div>
          <span className="text-[10px] text-slate-400">Shipped</span>
        </div>
        <div className="bg-white p-5 rounded-3xl border border-slate-200 shadow-sm space-y-1">
          <span className="text-xs font-semibold text-slate-500">COD Collected</span>
          <div className="text-2xl font-black text-[#DD0033]">৳ 348.5K</div>
          <span className="text-[10px] text-slate-400">Settled</span>
        </div>
        <div className="bg-white p-5 rounded-3xl border border-slate-200 shadow-sm space-y-1">
          <span className="text-xs font-semibold text-slate-500">Return Rate</span>
          <div className="text-2xl font-black text-amber-600">3.6%</div>
          <span className="text-[10px] text-slate-400">Low Risk</span>
        </div>
      </div>
    </div>
  );
}

