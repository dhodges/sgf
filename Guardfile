
def tempfile?(fname)
  File.basename(fname) =~ /^[.#]+.*/
end

guard :shell do
  watch /^[^.#]+.*.go$/ do |m|
    unless tempfile?(m[0])
      wd = Dir.getwd
      Dir.chdir './tests'
      puts `go test`
      Dir.chdir wd
    end
  end
end
