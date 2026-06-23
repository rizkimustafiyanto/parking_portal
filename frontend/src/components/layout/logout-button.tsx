"use client"

import { useRouter } from "next/navigation"
import { LogOutIcon } from "lucide-react"

import { buttonVariants } from "@/components/ui/button"
import { useAuthStore } from "@/features/auth"
import { cn } from "@/lib/utils"

function LogoutButton() {
  const router = useRouter()
  const logout = useAuthStore((state) => state.logout)

  function handleLogout() {
    logout()
    router.push("/login")
  }

  return (
    <button
      type="button"
      onClick={handleLogout}
      className={cn(buttonVariants({ variant: "outline" }), "w-full justify-start rounded-2xl px-4 py-6")}
    >
      <LogOutIcon className="size-4" />
      Logout
    </button>
  )
}

export { LogoutButton }
