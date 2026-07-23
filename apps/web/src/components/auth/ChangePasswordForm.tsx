'use client';

import { useState } from 'react';
import { useChangePasswordMutation } from '@/store/api/authApi';
import { useRouter } from 'next/navigation';
import Link from 'next/link';
import Image from 'next/image';
import { ArrowRight, Lock, AlertCircle } from 'lucide-react';
import { motion } from 'framer-motion';
import { Button } from '@/components/ui/Button';
import { Input } from '@/components/ui/Input';

export function ChangePasswordForm() {
  const [oldPassword, setOldPassword] = useState('');
  const [newPassword, setNewPassword] = useState('');
  const [errorMsg, setErrorMsg] = useState('');
  const [changePassword, { isLoading }] = useChangePasswordMutation();
  const router = useRouter();

  const handleSubmit = async (e: React.FormEvent) => {
    e.preventDefault();
    setErrorMsg('');
    try {
      const res = await changePassword({ oldPassword, newPassword }).unwrap();
      if (res.success) {
        router.push('/dashboard');
      }
    } catch (err: any) {
      setErrorMsg(err?.data?.message || 'Password update failed. Verify old password.');
    }
  };

  return (
    <div className="w-full max-w-md space-y-4">
      <motion.div
        initial={{ opacity: 0, scale: 0.95 }}
        animate={{ opacity: 1, scale: 1 }}
        className="w-full p-8 rounded-3xl bg-white border border-slate-200 shadow-xl space-y-6"
      >
        <div className="text-center space-y-3">
          <Link href="/" className="inline-block">
            <Image
              src="/logo.png"
              alt="Shipfide Logo"
              width={180}
              height={50}
              className="h-12 w-auto object-contain mx-auto"
              priority
            />
          </Link>
          <div>
            <h1 className="text-2xl font-extrabold tracking-tight text-[#404042]">Change Password</h1>
            <p className="text-xs text-slate-500 font-medium">Update security credentials to clear password reset flags</p>
          </div>
        </div>

        {errorMsg && (
          <div className="p-3 rounded-2xl bg-red-50 border border-red-200 text-[#DD0033] text-xs flex items-center gap-2">
            <AlertCircle className="w-4 h-4 shrink-0" />
            <span>{errorMsg}</span>
          </div>
        )}

        <form onSubmit={handleSubmit} className="space-y-4">
          <Input
            label="Current / Temporary Password"
            type="password"
            required
            value={oldPassword}
            onChange={(e) => setOldPassword(e.target.value)}
            placeholder="••••••••"
            leftIcon={<Lock className="w-4 h-4" />}
          />

          <Input
            label="New Password"
            type="password"
            required
            value={newPassword}
            onChange={(e) => setNewPassword(e.target.value)}
            placeholder="NewPassword@123"
            leftIcon={<Lock className="w-4 h-4" />}
          />

          <Button
            type="submit"
            isLoading={isLoading}
            rightIcon={<ArrowRight className="w-4 h-4" />}
            className="w-full py-3.5 mt-2"
          >
            Save New Password
          </Button>
        </form>

        <div className="text-center pt-2 border-t border-slate-100">
          <Link href="/dashboard" className="text-xs font-bold text-slate-500 hover:text-[#DD0033]">
            Back to Dashboard
          </Link>
        </div>
      </motion.div>

      {/* Manual Bottom Text */}
      <p className="text-center text-[11px] text-slate-500 font-medium">
        Shipfide • Delivering Trust, Every Time
      </p>
    </div>
  );
}
