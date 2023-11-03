all:
	@echo "do nothing"
	@echo "clean 规则 清空敏感信息公开到github中"

clean:

	git filter-branch --index-filter 'git rm -r --cached --ignore-unmatch imgdir' -- --all
	git push origin master --force

