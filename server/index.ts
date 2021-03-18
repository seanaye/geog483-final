import { connect } from "https://deno.land/x/redis/mod.ts";

const redis = await connect({
  hostname: "0.0.0.0",
  port: 6379,
});

const ok = await redis.set("hoge", "fuga");
const fuga = await redis.get("hoge");

console.log({ fuga })
