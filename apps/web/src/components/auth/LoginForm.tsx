'use client';

import { useState } from 'react';
import { useLoginMutation } from '@/store/api/authApi';
import { useRouter } from 'next/navigation';
import Link from 'next/link';
import Image from 'next/image';
import { ArrowRight, Lock, Mail, AlertCircle } from 'lucide-react';
import { motion } from 'framer-motion';
import { setAuthCookies } from '@/lib/cookies';
import { Button } from '@/components/ui/Button';
import { Input } from '@/components/ui/Input';

export function LoginForm() {
  const [email, setEmail] = useState('');
  const [password, setPassword] = useState('');
  const [errorMsg, setErrorMsg] = useState('');
  const [login, { isLoading }] = useLoginMutation();
  const router = useRouter();

  const handleSubmit = async (e: React.FormEvent) => {
    e.preventDefault();
    setErrorMsg('');
    try {
      const res = await login({ email, password }).unwrap();
      if (res.success) {
        if (res.data) {
          setAuthCookies(res.data.accessToken, res.data.refreshToken, res.data.sessionToken);
        }
        router.push('/dashboard');
      }
    } catch (err: any) {
      setErrorMsg(err?.data?.message || 'Login failed. Please check your credentials.');
    }
  };

  return (
    <div className="w-full max-w-md space-y-4">
      <motion.div
        initial={{ opacity: 0, scale: 0.95 }}
        animate={{ opacity: 1, scale: 1 }}
        className="w-full p-8 rounded-3xl bg-white border border-slate-200 shadow-xl space-y-6"
      >
        {/* Logo using /logo.png */}
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
            <h1 className="text-2xl font-extrabold tracking-tight text-[#404042]">Welcome Back</h1>
            <p className="text-xs text-slate-500 font-medium">Sign in to your Shipfide Courier account</p>
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
            label="Email Address"
            type="email"
            required
            value={email}
            onChange={(e) => setEmail(e.target.value)}
            placeholder="name@example.com"
            leftIcon={<Mail className="w-4 h-4" />}
          />

          <div className="space-y-1.5">
            <div className="flex items-center justify-between">
              <label className="text-xs font-semibold text-[#404042]">Password</label>
              <Link href="/auth/forgot-password" className="text-xs font-semibold text-[#DD0033] hover:underline">
                Forgot?
              </Link>
            </div>
            <Input
              type="password"
              required
              value={password}
              onChange={(e) => setPassword(e.target.value)}
              placeholder="••••••••"
              leftIcon={<Lock className="w-4 h-4" />}
            />
          </div>

          <Button
            type="submit"
            isLoading={isLoading}
            rightIcon={<ArrowRight className="w-4 h-4" />}
            className="w-full py-3.5 mt-2"
          >
            Sign In
          </Button>
        </form>

        <div className="text-center pt-2 border-t border-slate-100">
          <span className="text-xs text-slate-500">Don't have an account? </span>
          <Link href="/auth/register" className="text-xs font-bold text-[#DD0033] hover:underline">
            Register now
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
