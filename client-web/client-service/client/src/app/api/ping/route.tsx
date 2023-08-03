import { NextRequest, NextResponse } from "next/server";

export async function GET(req: NextRequest) {
    return new NextResponse(JSON.stringify({ message: "Pong" }), {
        status: 200,
      })
}