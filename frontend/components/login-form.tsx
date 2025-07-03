"use client"

import { useState } from "react"
import { useRouter } from "next/navigation"
import axios from "axios"
import { cn } from "@/lib/utils"
import { Button } from "@/components/ui/button"
import {
  Card,
  CardContent,
  CardDescription,
  CardHeader,
  CardTitle,
} from "@/components/ui/card"
import { Input } from "@/components/ui/input"
import { Label } from "@/components/ui/label"

export function LoginForm({
  className,
  ...props
}: React.ComponentProps<"div">) {
  const router = useRouter()

  const [formData, setFormData] = useState({
    sender_id: "",
    receiver_id: "",
    content: "",
  })

  const handleChange = (e: React.ChangeEvent<HTMLInputElement>) => {
    const { id, value } = e.target
    setFormData(prev => ({
      ...prev,
      [id]: value,
    }))
  }

  const handleSubmit = async (e: React.FormEvent) => {
    e.preventDefault()
    try {
      await axios.post("http://localhost:8080/api/messages", {
        sender_id: Number(formData.sender_id),
        receiver_id: Number(formData.receiver_id),
        content: formData.content,
      })
      router.push("/") // Redirect ke homepage
    } catch (error) {
      console.error("Failed to submit message:", error)
      alert("Gagal mengirim pesan")
    }
  }

  return (
    <div className={cn("flex flex-col gap-6", className)} {...props}>
      <Card>
        <CardHeader className="text-center">
          <CardTitle className="text-xl"></CardTitle>
          <CardDescription>
            Create a new message
          </CardDescription>
        </CardHeader>
        <CardContent>
          <form onSubmit={handleSubmit}>
            <div className="grid gap-6">
              <div className="grid gap-3">
                <Label htmlFor="sender_id">Sender ID</Label>
                <Input
                  id="sender_id"
                  type="number"
                  value={formData.sender_id}
                  onChange={handleChange}
                  required
                />
              </div>
              <div className="grid gap-3">
                <Label htmlFor="receiver_id">Receiver ID</Label>
                <Input
                  id="receiver_id"
                  type="number"
                  value={formData.receiver_id}
                  onChange={handleChange}
                  required
                />
              </div>
              <div className="grid gap-3">
                <Label htmlFor="content">Message</Label>
                <Input
                  id="content"
                  type="text"
                  value={formData.content}
                  onChange={handleChange}
                  required
                />
              </div>
              <Button type="submit" className="w-full">
                Submit
              </Button>
            </div>
          </form>
        </CardContent>
      </Card>
    </div>
  )
}
