
def tempfile(fname)
  fname = File.basename(fname)
  fname =~ /.*flymake.*/ || fname =~ /^[.#]+.*/
end

def runtest(fname)
  puts "#{File.basename(fname)} changed"
  wd = Dir.getwd

  Dir.chdir File.dirname(fname)
  `go test -v && go vet`
ensure
  Dir.chdir wd
end

guard :shell do
  watch /^[^.#]+.*.go$/ do |m|
    fname = m[0]
    runtest(fname) unless tempfile(fname)
  end
end
