defmodule StaticAssetServer do
  import Plug.Conn

  def init(options) do
    options
  end

  @spec call(Plug.Conn.t, any) :: Plug.Conn.t
  def call(conn, _opts) do
    case body(conn.request_path) do
      {:ok, response} ->
        conn
        |> resp(200, response)
        |> with_log()
        |> send_resp()
      {:error, :enoent} ->
        conn
        |> put_resp_content_type("text/plain")
        |> resp(404, "not found")
        |> with_log()
        |> send_resp()
      {:error, err} ->
        conn
        |> put_resp_content_type("text/plain")
        |> resp(500, to_string(err))
        |> with_log()
        |> send_resp()
    end
  end

  @spec body(String.t) :: {:ok, String.t} | {:error, atom}
  defp body(request_path) do
    # Yes, I know this is not safe and is a path traveral bug.
    path = String.replace_leading(request_path, "/", "./")
    if File.dir?(path) do
      dirlist(path)
    else
      File.read(path)
    end
  end

  @spec dirlist(String.t) :: {:ok, String.t} | {:error, any}
  defp dirlist(path) do
    case File.ls(path) do
      {:ok, files} ->
        header = "<!DOCTYPE html><html><head><title>Index for #{path}</title></head><body><ul>"
        footer = "</ul></body></html>"
        response = files
        |> Enum.map(fn file -> {file, Path.join(path, file)} end)
        |> Enum.map(fn {file, url} -> "<li><a href=\"#{url}\">#{file}</a></li>" end)
        |> Enum.reduce("", fn link, result -> result <> link end)
        response = header <> response <> footer
        {:ok, response}
      {:error, _} = err -> err
    end
  end

  @spec with_log(Plug.Conn.t) :: Plug.Conn.t
  defp with_log(conn) do
    log_line = Jason.encode!(%{
      "ts" => :os.system_time(:millisecond),
      "host" => conn.host,
      "method" => conn.method,
      "path" => conn.request_path,
      "remote_ip" => ip(conn.remote_ip),
      "status" => conn.status,
      "size" => String.length(conn.resp_body),
    })
    IO.puts(log_line)
    conn
  end

  @spec ip(tuple) :: String.t
  defp ip(remote_ip_tuple) do
    {a, b, c, d} = remote_ip_tuple
    Enum.join([a, b, c, d], ".")
  end
end
