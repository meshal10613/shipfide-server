
export default function SuperAdminSystemLogsPage() {
  return (
    <div className="space-y-6">
      <div>
        <h1 className="text-2xl font-black text-[#404042] tracking-tight">System Audit Logs</h1>
        <p className="text-xs text-slate-500">Live API event logs, DB migration traces, & security events</p>
      </div>

      <div className="bg-[#404042] text-slate-100 p-6 rounded-3xl font-mono text-xs space-y-2 shadow-inner overflow-x-auto">
        <p className="text-emerald-400">[2026-07-23 08:34:10] INFO Server started on: http://127.0.0.1:5000</p>
        <p className="text-blue-300">[2026-07-23 08:34:12] INFO Database auto-migrations completed successfully!</p>
        <p className="text-slate-300">[2026-07-23 08:34:15] POST /api/v1/auth/login 200 OK (User: superadmin.shipfide@gmail.com)</p>
        <p className="text-slate-300">[2026-07-23 08:34:18] GET /api/v1/shipments/track/SF-20260722-04829 200 OK</p>
      </div>
    </div>
  );
}

